package models

type UsersRequest struct {
	Name *string `json:"name"`
	Age  *int    `json:"age"`
}
