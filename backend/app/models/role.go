package models

import (
	"time"
)

type Role struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Role string `json:"role"`

	Users []User `gorm:"many2many:user_has_roles;" json:"users"`
}

type RoleHasPermissions struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	RoleID       uint      `json:"role_id"`
	PermissionID uint      `json:"permission_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	// DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at" swaggerignore:"true"`
}
