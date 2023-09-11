package entity

import (
	"errors"
	"fmt"
	"net/mail"
	"time"
	"unicode"

	"gorm.io/gorm"
)

const (
	ROLE_DEFAULT = "default"
)

type User struct {
	gorm.Model
	Name     string `gorm:"size:200"`
	Email    string `gorm:"unique; not null"`
	Role     string `gorm:"not null"`
	Password string `gorm:"not null"`
}

// type UserDTO struct {
// 	ID    int    `json:"ID"`
// 	Name  string `json:"name"`
// 	Email string `json:"email"`
// 	Role  string `json:"role"`
// }

type UserOutput struct {
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Role      string    `json:"role,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToDTO(user *User) *UserOutput {
	return &UserOutput{
		ID:        int(user.ID),
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func NewUser(name, email, password string) (*User, error) {
	user := &User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	if err := user.Validate(); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) Validate() error {
	if u.Name == "" {
		return errors.New("name is empty")
	}
	_, err := mail.ParseAddress(u.Email)
	if err != nil {
		return fmt.Errorf("einvalid email: %v", err)
	}

	if !isPasswordValid(u.Password) {
		return errors.New("invalid password")
	}
	return nil
}

func isPasswordValid(p string) bool {
	var (
		hasNumber  = false
		hasUpper   = false
		hasLower   = false
		hasSpecial = false
	)
	if len(p) < 7 {
		return false
	}

	for _, c := range p {
		switch {
		case unicode.IsNumber(c):
			hasNumber = true
		case unicode.IsUpper(c):
			hasUpper = true
		case unicode.IsLower(c):
			hasLower = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			hasSpecial = true
		}
	}
	return hasNumber && hasUpper && hasLower && hasSpecial
}
