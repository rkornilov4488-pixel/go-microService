package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"myMicroService/internal/models"
)

type OrdersGenerator struct{}

func (og *OrdersGenerator) GetRespOnOrderCreate(ctx context.Context, request *models.OrdersRequest) (models.OrderResponse, error) {

	if ctx.Err() != nil {
		return models.OrderResponse{}, ctx.Err()
	}

	if request == nil {
		return models.OrderResponse{Error: "request body required"}, errors.New("request body is empty")
	}
	if request.UserId == nil || *request.UserId == "" {
		return models.OrderResponse{Error: "user_id is required"}, errors.New("userId is empty")
	}
	if request.Items == nil || len(request.Items) < 1 {
		return models.OrderResponse{Error: "items required is not empty"}, errors.New("items is empty")
	}
	for _, item := range request.Items {
		if item.Product == nil || *item.Product == "" {
			return models.OrderResponse{Error: "product is required"}, errors.New("product is empty")
		}
		if item.Price == nil || *item.Price <= 0 {
			return models.OrderResponse{Error: "price must be greater than 0"}, errors.New("price is empty")
		}
		if item.Quantity == nil || *item.Quantity <= 0 {
			return models.OrderResponse{Error: "quantity must be greater than 0"}, errors.New("quantity is empty")
		}
	}

	var sum int
	items := request.Items
	for _, item := range items {
		sum += *item.Price * *item.Quantity
	}

	return models.OrderResponse{
		OrderId: uuid.New().String(), UserId: *request.UserId, Total: sum,
	}, nil
}
