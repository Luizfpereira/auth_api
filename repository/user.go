package repository

import (
	"auth_api/entity"

	"gorm.io/gorm"
)

type UserRepositoryPSQL struct {
	Instance *gorm.DB
}

func NewUserRepositoryPSQL(instance *gorm.DB) *UserRepositoryPSQL {
	return &UserRepositoryPSQL{Instance: instance}
}

func (u *UserRepositoryPSQL) CreateUser(user *entity.User) (*entity.UserOutput, error) {
	res := u.Instance.Create(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	return entity.ToDTO(user), nil
}
