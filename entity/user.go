package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"size:200"`
	Email    string `gorm:"unique; not null"`
	Role     string `gorm:"not null"`
	Password string `gorm:"not null"`
}
