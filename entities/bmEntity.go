package entities

import "time"

type WhistBm struct {
	ID          int       `gorm:"size:11;primary_key;not null;auto_increment" json:"id"`
	Title       string    `gorm:"size:255;null;" json:"title"`
	ProductName string    `gorm:"size:255;null" json:"product_name"`
	Description string    `gorm:"type:text;null;" json:"description"`
	Location    string    `gorm:"type:text;null" json:"location"`
	DueDate     time.Time `gorm:"null;" json:"due_date"`
	Status      string    `gorm:"size:255;null" json:"status"`
	IsDeleted   bool      `gorm:"default:false;not null;" json:"is_deleted"`
}
