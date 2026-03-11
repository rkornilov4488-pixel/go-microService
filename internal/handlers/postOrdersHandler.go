package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"myMicroService/internal/models"
	"net/http"
)

type orderHandlerInterface interface {
	GetRespOnOrderCreate(ctx context.Context, request *models.OrdersRequest) (models.OrderResponse, error)
}

type CreateOrderHandler struct {
	Service orderHandlerInterface
}

func return400ErrResponse(w http.ResponseWriter, err error) {
	log.Println(err)
	w.WriteHeader(http.StatusBadRequest)
}

func (h *CreateOrderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	var requestStruct models.OrdersRequest
	decodeReqBodyToStruct := json.NewDecoder(r.Body).Decode(&requestStruct)
	if decodeReqBodyToStruct != nil {
		return400ErrResponse(w, decodeReqBodyToStruct)
		return
	}
	respStruct, serviceErr := h.Service.GetRespOnOrderCreate(r.Context(), &requestStruct)
	if serviceErr != nil {
		if errors.Is(serviceErr, context.Canceled) || errors.Is(serviceErr, context.DeadlineExceeded) {
			return
		}
		byteResp, marshalToBytesErr := json.Marshal(respStruct)
		if marshalToBytesErr != nil {
			return400ErrResponse(w, marshalToBytesErr)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		_, writeRespBodyErr := w.Write(byteResp)
		if writeRespBodyErr != nil {
			return400ErrResponse(w, writeRespBodyErr)
			return
		}
		return
	} else {
		byteResp, marshalToBytesErr := json.Marshal(respStruct)
		if marshalToBytesErr != nil {
			return400ErrResponse(w, marshalToBytesErr)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, writeRespBodyErr := w.Write(byteResp)
		if writeRespBodyErr != nil {
			return400ErrResponse(w, writeRespBodyErr)
			return
		}
		return
	}
}
