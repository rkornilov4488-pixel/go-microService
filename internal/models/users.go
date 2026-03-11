package models

import "time"

type UsersRequest struct {
	UserId   *string `json:"user_id"`
	Name     *string `json:"name"`
	LastName *string `json:"last_name"`
	Surname  *string `json:"surname"`
	Age      *int    `json:"age"`
}

type UsersResponse struct {
	Error string `json:"error,omitempty"`
	Id    string `json:"id,omitempty"`
}

type User struct {
	Id        string
	Name      string
	Age       int
	UserId    string
	LastName  string
	Surname   string
	CreatedAt time.Time
}
