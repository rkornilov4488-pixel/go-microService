package models

type UsersRequest struct {
	Name *string `json:"name"`
	Age  *int    `json:"age"`
}

type UsersResponse struct {
	Error string `json:"error,omitempty"`
	Id    string `json:"id,omitempty"`
}

type User struct {
	Id   string
	Name string
	Age  int
}
