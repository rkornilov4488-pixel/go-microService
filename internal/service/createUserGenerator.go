package service

import (
	"context"
	"github.com/goliatone/hashid/pkg/hashid"
	"myMicroService/internal/models"
	"myMicroService/internal/repository"
	"strconv"
)

type UserService struct {
	UserRepo repository.UserRepo
}

func (us *UserService) GetUserHashId(ctx context.Context, request models.UsersRequest) (models.UsersResponse, error) {
	if ctxErr := ctx.Err(); ctxErr != nil {
		return models.UsersResponse{}, ctxErr
	} else {
		id, err := hashid.New(*request.Name + strconv.Itoa(*request.Age))
		insertErr := us.UserRepo.InsertUser(ctx, models.User{Id: id, Name: *request.Name, Age: *request.Age})
		if insertErr != nil {
			return models.UsersResponse{}, insertErr
		}
		return models.UsersResponse{Id: id}, err
	}
}
