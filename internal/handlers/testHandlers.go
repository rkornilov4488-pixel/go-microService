package handlers

import (
	"encoding/json"
	"log"
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
