package models

import "time"

type Product struct {
	ID            uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Category_id   uint32 `json:"category_id"`
	Category      Category
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Status        int       `json:"status"`
	Image         string    `json:"image"`
	Image_slide   string    `json:"image_slide"`
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	ProductDetail Product_details
	Rating        []Rating
}

func (p *Product) TableName() string {
	return "products"
}
