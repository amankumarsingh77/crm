package models

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

type User struct {
	UserID      uuid.UUID  `json:"user_id" db:"user_id" redis:"user_id" validate:"omitempty"`
	Email       string     `json:"email,omitempty" db:"email" redis:"email" validate:"required,email,lte=60"`
	Password    string     `json:"password,omitempty" db:"password" redis:"password" validate:"required,min=8"`
	FirstName   string     `json:"first_name" db:"first_name" redis:"first_name" validate:"required,lte=30"`
	LastName    string     `json:"last_name" db:"last_name" redis:"last_name" validate:"required,lte=30"`
	Role        *string    `json:"role,omitempty" db:"role" redis:"role" validate:"required,oneof=admin user,lte=10"`
	Phone       *string    `json:"phone_number,omitempty" db:"phone_number" redis:"phone_number" validate:"required,e164,lte=20"`
	IsActive    bool       `json:"is_active" db:"is_active" redis:"is_active"`
	LastLoginAt *time.Time `json:"last_login_at,omitempty" db:"last_login_at" redis:"last_login_at"`
	CreatedAt   time.Time  `json:"created_at,omitempty" db:"created_at" redis:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at,omitempty" db:"updated_at" redis:"updated_at"`
}

func (u *User) Sanitize() {
	u.Password = ""
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) ComparePassword(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return err
	}
	return nil
}

func (u *User) SanitizePassword() {
	u.Password = ""
}

func (u *User) PrepareCreate() error {
	u.Email = strings.ToLower(strings.TrimSpace(u.Email))
	u.Password = strings.TrimSpace(u.Password)
	if err := u.HashPassword(); err != nil {
		return err
	}
	if u.Phone != nil {
		*u.Phone = strings.TrimSpace(*u.Phone)
	}
	if u.Role != nil {
		*u.Role = strings.ToLower(strings.TrimSpace(*u.Role))
	}
	return nil
}

type UsersList struct {
	TotalCount int     `json:"total_count" `
	TotalPages int     `json:"total_pages" `
	Page       int     `json:"page" `
	Size       int     `json:"size" `
	Users      []*User `json:"users" `
}

type UserWithToken struct {
	User  *User  `json:"user"`
	Token string `json:"token"`
}
