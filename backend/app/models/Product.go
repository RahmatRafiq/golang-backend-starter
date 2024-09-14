package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Reference   string         `gorm:"unique" json:"reference"`
	StoreID     uint           `json:"store_id"`
	CategoryID  uint           `json:"category_id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Price       float64        `json:"price"`
	Margin      float64        `json:"margin"`
	Stock       int            `json:"stock"`
	Sold        int            `json:"sold"`
	Images      string         `json:"images"` // Assume JSON stored as string
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at" swaggerignore:"true"`
}
