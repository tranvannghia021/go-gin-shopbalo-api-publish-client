package models

import "time"

type StatusRating string
type Rating struct {
	ID          uint32       `gorm:"primary_key;auto_increment" json:"id"`
	Customer_id int          `json:"customer_id"`
	Product_id  int          `json:"product_id"`
	Point       float32      `json:"point"`
	Point_avg   float64      `json:"point_avg"`
	Status      StatusRating `gorm:"type:enum('pending', 'pushlished')" json:"status"`
	Content     string       `json:"content"`
	Image       string       `json:"image"`
	CreatedAt   time.Time    `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time    `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (r *Rating) TableName() string {
	return "ratings"
}
