package main

import (
	"myMicroService/internal/handlers"
	"myMicroService/internal/service"
	"net/http"
)

func main() {
	userService := &service.UserService{}
	handler := handlers.CreateUserHandler{Service: userService}
	http.Handle("/users", &handler)
	http.ListenAndServe("localhost:8080", nil)
}
