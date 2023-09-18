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

func (u *UserRepositoryPSQL) GetUserByEmail(email string) (*entity.User, error) {
	var user *entity.User
	res := u.Instance.Where("email = ?", email).First(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}

func (u *UserRepositoryPSQL) GetUserById(id int) (*entity.UserOutput, error) {
	var user *entity.User
	res := u.Instance.First(&user, id)
	if res.Error != nil {
		return nil, res.Error
	}
	return entity.ToDTO(user), nil
}
