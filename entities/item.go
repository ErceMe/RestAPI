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
	Created_at  time.Time `gorm:"autoCreateTime" json:"created_at"`
	Updated_at  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type Order struct {
	Id            uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Customer_name string    `gorm:"type:varchar(45)" json:"customer_name"`
	Ordered_at    time.Time `gorm:"type:datetime" json:"ordered_at"`
	Created_at    time.Time `gorm:"AutoCreateTime" json:"created_at"`
	Updated_at    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Items         []Item    `json:"items"`
}
