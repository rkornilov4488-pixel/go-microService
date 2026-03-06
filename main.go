package main

import (
	"myMicroService/internal/handlers"
	"myMicroService/internal/middleWare"
	"myMicroService/internal/service"
	"net/http"
)

func main() {
	userService := &service.UserService{}
	createUserHandler := handlers.CreateUserHandler{Service: userService}

	orderGenerator := &service.OrdersGenerator{}
	createOrderHandler := handlers.CreateOrderHandler{Service: orderGenerator}

	loggingUserHandler := middleWare.LoggingMiddleware(&createUserHandler)
	loggingOrderHandler := middleWare.LoggingMiddleware(&createOrderHandler)

	http.Handle("/users", loggingUserHandler)
	http.Handle("/orders", loggingOrderHandler)
	http.ListenAndServe("localhost:8080", nil)
}
