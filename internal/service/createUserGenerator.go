package service

import (
	"github.com/goliatone/hashid/pkg/hashid"
	"myMicroService/internal/models"
	"strconv"
)

type UserService struct{}

func (us *UserService) GetUserHashId(request models.UsersRequest) (models.UsersResponse, error) {
	id, err := hashid.New(*request.Name + strconv.Itoa(*request.Age))
	return models.UsersResponse{Id: id}, err
}
