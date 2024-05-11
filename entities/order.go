package entities

import "time"

type Order struct {
	Id        string    `json:"product"`
	UserId    string    `json:"user_id"`
	Total     float64   `json:"total"`
	Status    string    `json:"status"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
}

type OrderItem struct {
	Id        string    `json:"id"`
	OrderId   string    `json:"order_id"`
	ProductId string    `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}
