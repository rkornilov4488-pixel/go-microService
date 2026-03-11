package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"myMicroService/internal/models"
	"myMicroService/internal/repository"
)

type UserService struct {
	UserRepo repository.UserRepo
}

func (us *UserService) CreateUser(ctx context.Context, request *models.UsersRequest) (models.UsersResponse, error) {
	if ctxErr := ctx.Err(); ctxErr != nil {
		return models.UsersResponse{}, ctxErr
	}

	if request == nil {
		return models.UsersResponse{Error: "request body is empty"}, errors.New("request body is empty")
	}
	if request.UserId == nil || *request.UserId == "" {
		return models.UsersResponse{Error: "user_id is empty"}, errors.New("user_id is empty")
	}
	if request.Name == nil || *request.Name == "" {
		return models.UsersResponse{Error: "name is empty"}, errors.New("name is empty")
	}
	if request.LastName == nil || *request.LastName == "" {
		return models.UsersResponse{Error: "last_name is empty"}, errors.New("last_name is empty")
	}
	if request.Surname == nil || *request.Surname == "" {
		return models.UsersResponse{Error: "surname is empty"}, errors.New("surname is empty")
	}
	if request.Age == nil {
		return models.UsersResponse{Error: "age is empty"}, errors.New("age is empty")
	}
	if *request.Age < 14 {
		return models.UsersResponse{Error: "age is less than 14"}, errors.New("age is less than 14")
	}

	id := uuid.New().String()
	insertErr := us.UserRepo.InsertUser(
		ctx,
		models.User{
			Id:       id,
			Name:     *request.Name,
			Age:      *request.Age,
			LastName: *request.LastName,
			Surname:  *request.Surname,
			UserId:   *request.UserId,
		},
	)
	if insertErr != nil {
		return models.UsersResponse{}, insertErr
	}
	return models.UsersResponse{Id: id}, nil
}
