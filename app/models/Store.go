package models

import (
	"time"

	"gorm.io/gorm"
)

type Store struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `json:"name"`
	Phone     string         `json:"phone"`
	Address   string         `json:"address"`
	City      string         `json:"city"`
	State     string         `json:"state"`
	Country   string         `json:"country"`
	Zip       string         `json:"zip"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at" swaggerignore:"true"`
}
