package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"myMicroService/internal/models"
	"net/http"
)

func validateAndReturnNegativeUsersResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	negativeResp, _ := json.Marshal(models.UsersResponse{Error: "incorrect params"})
	w.WriteHeader(http.StatusBadRequest)
	_, respErr := w.Write(negativeResp)
	if respErr != nil {
		log.Println(respErr)
	}
}

type CreateUserHandler struct {
	Service createUserInterface
}

type createUserInterface interface {
	GetUserHashId(ctx context.Context, request models.UsersRequest) (models.UsersResponse, error)
}

func (h *CreateUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()
	if r.Body == nil {
		validateAndReturnNegativeUsersResponse(w)
		return
	}
	var structReq models.UsersRequest
	err := json.NewDecoder(r.Body).Decode(&structReq)
	if err != nil {
		validateAndReturnNegativeUsersResponse(w)
		return
	}
	if structReq.Name == nil || structReq.Age == nil || *structReq.Name == "" {
		validateAndReturnNegativeUsersResponse(w)
		return
	}

	structResp, serviceErr := h.Service.GetUserHashId(r.Context(), structReq)
	if serviceErr != nil {
		if errors.Is(serviceErr, context.Canceled) || errors.Is(serviceErr, context.DeadlineExceeded) {
			return
		} else {
			validateAndReturnNegativeUsersResponse(w)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	parsedResp, _ := json.Marshal(structResp)
	_, errResp := w.Write(parsedResp)
	if errResp != nil {
		log.Println(errResp)
	}
}
