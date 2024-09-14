package models

import (
	"time"

	"gorm.io/gorm"
)

type Permission struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	Permission string `json:"permission"`
}

type UserHasPermissions struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	UserID       uint           `json:"user_id"`
	PermissionID uint           `json:"permission_id"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at" swaggerignore:"true"`
}
