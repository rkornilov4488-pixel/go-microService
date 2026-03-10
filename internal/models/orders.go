package models

type OrderItems struct {
	Product  *string `json:"product"`
	Price    *int    `json:"price"`
	Quantity *int    `json:"quantity"`
}

type OrdersRequest struct {
	UserId *string      `json:"user_id"`
	Items  []OrderItems `json:"items"`
}

type OrderResponse struct {
	OrderId string `json:"order_id,omitempty"`
	UserId  string `json:"user_id,omitempty"`
	Total   int    `json:"total,omitempty"`
	Error   string `json:"error,omitempty"`
}
