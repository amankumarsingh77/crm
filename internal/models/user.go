package models

import (
	"github.com/gofrs/uuid"
	"time"
)

type User struct {
	ID          uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" redis:"id"`
	Email       string    `json:"email" gorm:"uniqueIndex;not null" redis:"email" validate:"required,email"`
	Password    string    `json:"-" gorm:"not null" validate:"required,min=8"`
	FirstName   string    `json:"first_name" gorm:"not null" redis:"first_name" validate:"required"`
	LastName    string    `json:"last_name" gorm:"not null" redis:"last_name" validate:"required"`
	Role        *string   `json:"role" gorm:"type:user_role;not null" redis:"role" validate:"required,oneof=admin user"`
	Phone       *string   `json:"phone" gorm:"uniqueIndex" redis:"phone" validate:"required,e164"`
	IsActive    bool      `json:"is_active" gorm:"default:true" redis:"is_active"`
	LastLoginAt time.Time `json:"last_login_at,omitempty" gorm:"index" redis:"last_login_at"`
	CreatedAt   time.Time `json:"created_at" gorm:"not null" redis:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"not null" redis:"updated_at"`
}

type UsersList struct {
	TotalCount int     `json:"total_count" `
	TotalPages int     `json:"total_pages" `
	Page       int     `json:"page" `
	Size       int     `json:"size" `
	Users      []*User `json:"users" `
}

type UserWithToken struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}
