package main

import (
	"myMicroService/internal/handlers"
	"myMicroService/internal/service"
	"net/http"
)

func main() {
	userService := &service.UserService{}
	handler := handlers.MyHandlers{Service: userService}
	http.Handle("/users", &handler)
	http.ListenAndServe("localhost:8080", nil)
}
