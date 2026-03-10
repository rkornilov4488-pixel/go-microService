package repository

import (
	"context"
	"database/sql"
	"log"
	"myMicroService/internal/models"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (repo *UserRepo) InsertUser(ctx context.Context, user models.User) error {
	_, insertErr := repo.db.ExecContext(ctx, "INSERT INTO users (id, name, age) VALUES ($1, $2, $3)", user.Id, user.Name, user.Age)
	if insertErr != nil {
		log.Println(insertErr)
		return insertErr
	}
	return nil
}
