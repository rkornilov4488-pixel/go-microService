package handlers

import (
	"encoding/json"
	"log"
	"myMicroService/internal/models"
	"net/http"
)

type orderHandlerInterface interface {
	GetRespOnOrderCreate(request *models.OrdersRequest) (models.OrderResponse, error)
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
	defer r.Body.Close()
	var requestStruct models.OrdersRequest
	decodeReqBodyToStruct := json.NewDecoder(r.Body).Decode(&requestStruct)
	if decodeReqBodyToStruct != nil {
		return400ErrResponse(w, decodeReqBodyToStruct)
		return
	}
	respStruct, errValidationReq := h.Service.GetRespOnOrderCreate(&requestStruct)
	if errValidationReq != nil {
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
