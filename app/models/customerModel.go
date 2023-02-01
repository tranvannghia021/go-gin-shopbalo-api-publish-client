package models

import "time"

type Customer struct {
	ID           uint32    `gorm:"primary_key;auto_increment" json:"id"`
	First_name   int       `json:"first_name"`
	Last_name    int       `json:"Last_name"`
	Gender       string    `json:"gender"`
	Phone        string    `json:"phone"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	Point        int       `json:"point"`
	Avatar       string    `json:"avatar"`
	Status       int       `json:"status"`
	Address      string    `json:"address"`
	Created_date time.Time `json:"created_date"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (c *Customer) TableName() string {
	return "customers"
}
