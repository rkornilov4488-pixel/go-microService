package service

import (
	"context"
	"github.com/goliatone/hashid/pkg/hashid"
	"myMicroService/internal/models"
	"strconv"
)

type UserService struct{}

func (us *UserService) GetUserHashId(ctx context.Context, request models.UsersRequest) (models.UsersResponse, error) {
	if ctxErr := ctx.Err(); ctxErr != nil {
		return models.UsersResponse{}, ctxErr
	} else {
		id, err := hashid.New(*request.Name + strconv.Itoa(*request.Age))
		return models.UsersResponse{Id: id}, err
	}
}
