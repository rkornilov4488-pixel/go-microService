package repository

import (
	"context"
	"database/sql"
	"fmt"
	"myMicroService/internal/models"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (repo *UserRepo) InsertUser(ctx context.Context, user models.User) error {
	_, insertErr := repo.db.ExecContext(
		ctx,
		"INSERT INTO users (id, name, age, user_id, last_name, surname) VALUES ($1, $2, $3, $4, $5, $6)",
		user.Id, user.Name, user.Age, user.UserId, user.LastName, user.Surname,
	)
	if insertErr != nil {
		return fmt.Errorf("error inserting user: %w", insertErr)
	}
	return nil
}
