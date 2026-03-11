package main

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
	"myMicroService/internal/handlers"
	"myMicroService/internal/middleWare"
	"myMicroService/internal/repository"
	"myMicroService/internal/service"
	"net/http"
)

func main() {

	db, openConErr := sql.Open("pgx", "postgres://romankornilov@localhost:5432/test_db?sslmode=disable")
	if openConErr != nil {
		log.Fatal(openConErr)
	}
	pingDBErr := db.Ping()
	if pingDBErr != nil {
		log.Fatal(pingDBErr)
	}
	log.Println("db connected")

	ur := repository.NewUserRepo(db)

	userService := &service.UserService{UserRepo: *ur}
	createUserHandler := handlers.CreateUserHandler{Service: userService}

	orderGenerator := &service.OrdersGenerator{}
	createOrderHandler := handlers.CreateOrderHandler{Service: orderGenerator}

	loggingUserHandler := middleWare.LoggingMiddleware(&createUserHandler)
	loggingOrderHandler := middleWare.LoggingMiddleware(&createOrderHandler)

	http.Handle("/users", loggingUserHandler)
	http.Handle("/orders", loggingOrderHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
