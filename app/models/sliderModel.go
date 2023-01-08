package models

import "time"

type Slider_status string
type Slider struct {
	ID           uint32        `gorm:"primary_key;auto_increment" json:"id"`
	Name         string        `gorm:"size:255;not null;unique" json:"name"`
	Description  string        `gorm:"size:255;not null;unique" json:"description"`
	Image_slider string        `gorm:"size:255;not null;unique" json:"image"`
	Status       Slider_status `gorm:"type:enum('Active', 'InActive')" json:"status"`
	Url          string        `gorm:"size:255;not null;unique" json:"url"`
	CreatedAt    time.Time     `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time     `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
