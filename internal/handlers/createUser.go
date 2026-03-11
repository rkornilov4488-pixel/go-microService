package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
	"log"
	"myMicroService/internal/constants"
	"myMicroService/internal/models"
	"net/http"
)

func validateAndReturnNegativeUsersResponse(w http.ResponseWriter, r *models.UsersResponse) {
	var structResp models.UsersResponse
	if r == nil {
		structResp = models.UsersResponse{Error: "incorrect params"}
	} else {
		structResp = *r
	}
	negativeResp, _ := json.Marshal(structResp)
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
	CreateUser(ctx context.Context, request *models.UsersRequest) (models.UsersResponse, error)
}

func (h *CreateUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	defer func() {
		closeApiErr := r.Body.Close()
		if closeApiErr != nil {
			log.Println(closeApiErr)
		}
	}()

	var structReq models.UsersRequest
	err := json.NewDecoder(r.Body).Decode(&structReq)
	if err != nil {
		validateAndReturnNegativeUsersResponse(w, nil)
		return
	}

	var pgErr *pgconn.PgError
	structResp, serviceErr := h.Service.CreateUser(r.Context(), &structReq)
	if serviceErr != nil {
		if errors.Is(serviceErr, context.Canceled) || errors.Is(serviceErr, context.DeadlineExceeded) {
			return
		} else if errors.As(serviceErr, &pgErr) && pgErr.Code == constants.PgUniqViolation {
			w.WriteHeader(http.StatusConflict)
			return
		} else {
			validateAndReturnNegativeUsersResponse(w, &structResp)
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
