package models

import "time"

type Category struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:255;not null;unique" json:"name"`
	Parent_id uint32    `json:"parent_id"`
	Image     string    `gorm:"size:255;not null;unique" json:"image"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (c *Category) TableName() string {
	return "categories"
}
