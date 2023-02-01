package models

import "time"

type Product_details struct {
	ID         uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Product_id uint32    `json:"product_id"`
	Code_color string    `gorm:"size:255;" json:"code_color"`
	Amount     int       `json:"amount"`
	Price      int       `json:"price"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Product_details) TableName() string {
	return "product_details"
}
