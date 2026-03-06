package models

type UsersResponse struct {
	Error string `json:"error,omitempty"`
	Id    string `json:"id,omitempty"`
}

type OrderResponse struct {
	OrderId string `json:"order_id,omitempty"`
	UserId  string `json:"user_id,omitempty"`
	Total   int    `json:"total,omitempty"`
	Error   string `json:"error,omitempty"`
}
