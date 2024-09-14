package models

import (
	"time"

	"github.com/RahmatRafiq/golang_backend_starter/app/helpers"
	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Reference string         `json:"reference"`
	Username  string         `json:"username"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	JwtToken  string         `json:"jwt_token" swaggerignore:"true"`
	FcmToken  string         `json:"fcm_token" swaggerignore:"true"`
	Pin       string         `json:"pin"`
	CreatedAt time.Time      `json:"created_at" swaggerignore:"true"`
	UpdatedAt time.Time      `json:"updated_at" swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at" swaggerignore:"true"`

	Roles []Role `gorm:"many2many:user_has_roles;" json:"roles" swaggerignore:"true"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	reference := helpers.GenerateReference("USR")
	password, err := helpers.HashPasswordArgon2(u.Password, helpers.DefaultParams)
	if err != nil {
		println(err.Error())
		return
	}
	// pin, err := helpers.HashPasswordBcrypt(u.Pin)
	pin, err := helpers.HashPasswordArgon2(u.Pin, helpers.DefaultParams)
	if err != nil {
		println(err.Error())
	}
	tx.Statement.SetColumn("reference", reference)
	tx.Statement.SetColumn("password", password)
	tx.Statement.SetColumn("pin", pin)

	return
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type UserHasRoles struct {
	ID     uint `gorm:"primaryKey" json:"id"`
	UserID uint `json:"user_id"`
	RoleID uint `json:"role_id"`
}
