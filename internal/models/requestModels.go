package models

type UsersRequest struct {
	Name *string `json:"name"`
	Age  *int    `json:"age"`
}

type OrderItems struct {
	Product  *string `json:"product"`
	Price    *int    `json:"price"`
	Quantity *int    `json:"quantity"`
}

type OrdersRequest struct {
	UserId *string      `json:"user_id"`
	Items  []OrderItems `json:"items"`
}
