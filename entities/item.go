package entities

import (
	"time"
)

type Item struct {
	Id_item     int       `json:"id_item"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Quantity    int       `json:"quantity"`
	Order_id    int       `json:"order_id"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}

type Order struct {
	Id            int       `json:"id"`
	Costumer_name string    `json:"costumer_name"`
	Items         []Item    `json:"items"`
	Orderer_at    time.Time `json:"ordered_at"`
	Created_at    time.Time `json:"created_at"`
	Updated_at    time.Time `json:"updated_at"`
}
