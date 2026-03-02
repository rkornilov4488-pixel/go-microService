package handlers

import (
	"encoding/json"
	"log"
	"myMicroService/internal/models"
	"myMicroService/internal/service"
	"net/http"
)

type SuccessHealthResponse struct {
	Message string `json:"message"`
}

type GreetResponse struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(SuccessHealthResponse{"ok"})
	if err != nil {
		panic(err)
	}
	_, errWrite := w.Write(resp)
	if errWrite != nil {
		panic(errWrite)
	}
}

func GreetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	queryParamName := r.URL.Query().Get("name")
	if queryParamName == "" {
		w.WriteHeader(http.StatusBadRequest)
		resp, err := json.Marshal(GreetResponse{Error: "name is required"})
		if err != nil {
			log.Fatal(err)
		} else {
			_, err := w.Write(resp)
			if err != nil {
				log.Fatal(err)
			}
		}
	} else {
		resp, err := json.Marshal(GreetResponse{Message: "Hello, " + queryParamName + "!"})
		if err != nil {
			log.Fatal(err)
		} else {
			w.WriteHeader(http.StatusOK)
			_, err := w.Write(resp)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func validateAndReturnNegativeUsersResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	negativeResp, _ := json.Marshal(models.UsersResponse{Error: "incorrect params"})
	w.WriteHeader(http.StatusBadRequest)
	_, respErr := w.Write(negativeResp)
	if respErr != nil {
		log.Println(respErr)
	}
}

type MyHandlers struct {
	Service *service.UserService
}

func (h *MyHandlers) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	requestBody := r.Body
	defer r.Body.Close()
	if requestBody == nil {
		validateAndReturnNegativeUsersResponse(w)
		return
	}
	var structReq models.UsersRequest
	err := json.NewDecoder(requestBody).Decode(&structReq)
	if err != nil {
		validateAndReturnNegativeUsersResponse(w)
		return
	}
	if structReq.Name == nil || structReq.Age == nil || *structReq.Name == "" {
		validateAndReturnNegativeUsersResponse(w)
		return
	}

	structResp, parsErr := h.Service.CreateUser(structReq)
	if parsErr != nil {
		validateAndReturnNegativeUsersResponse(w)
		return
	}
	w.WriteHeader(http.StatusOK)
	parsedResp, _ := json.Marshal(structResp)
	_, errResp := w.Write(parsedResp)
	if errResp != nil {
		log.Println(errResp)
	}
}
