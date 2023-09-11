package gateway

import "auth_api/entity"

type UserGateway interface {
	CreateUser(user *entity.User) (*entity.UserOutput, error)
}
