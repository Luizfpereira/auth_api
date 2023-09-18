package gateway

import "auth_api/entity"

type UserGateway interface {
	CreateUser(user *entity.User) (*entity.UserOutput, error)
	GetUserByEmail(email string) (*entity.User, error)
	GetUserById(id int) (*entity.UserOutput, error)
}
