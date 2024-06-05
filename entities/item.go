package entities

import (
	"time"
)

type Item struct {
	Id          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"type:varchar(45)" json:"name"`
	Description string    `gorm:"type:varchar(100)" json:"description"`
	Quantity    int       `gorm:"type:int" json:"quantity"`
	Order_id    uint      `form:"order_id" json:"order_id"`
	Created_at  time.Time `gorm:"type:datetime" json:"created_at"`
	Updated_at  time.Time `gorm:"type:timestamp(5)" json:"updated_at"`
}

type Order struct {
	Id            uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Costumer_name string    `gorm:"type:varchar(45)" json:"costumer_name"`
	Orderer_at    time.Time `gorm:"type:datetime" json:"orderer_at"`
	Created_at    time.Time `gorm:"type:datetime" json:"created_at"`
	Updated_at    time.Time `gorm:"type:timestamp" json:"updated_at"`
	Items         []Item    `json:"items"`
}
