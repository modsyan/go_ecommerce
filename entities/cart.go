package entities

import "time"

type Cart struct {
	Id        string    `json:"id"`
	UserId    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type CartItem struct {
	ProductId string `json:"product_id"`
	CartId    string `json:"cart_id"`
}
