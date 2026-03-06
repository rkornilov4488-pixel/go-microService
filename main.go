package main

import (
	"myMicroService/internal/handlers"
	"myMicroService/internal/service"
	"net/http"
)

func main() {
	userService := &service.UserService{}
	createUserHandler := handlers.CreateUserHandler{Service: userService}
	orderGenerator := &service.OrdersGenerator{}
	createOrderHandler := handlers.CreateOrderHandler{Service: orderGenerator}
	http.Handle("/users", &createUserHandler)
	http.Handle("/orders", &createOrderHandler)
	http.ListenAndServe("localhost:8080", nil)
}
