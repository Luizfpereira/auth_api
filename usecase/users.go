package usecase

import (
	"auth_api/entity"
	"auth_api/gateway"
	"errors"
)

type UserUsecase struct {
	gateway gateway.UserGateway
}

type UserInput struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
}

func NewUserUsecase(gateway gateway.UserGateway) *UserUsecase {
	return &UserUsecase{gateway: gateway}
}

func (u *UserUsecase) CreateUser(userInput UserInput) (*entity.UserOutput, error) {
	if userInput.Password != userInput.PasswordConfirm {
		return nil, errors.New("password and confirmation password don't match")
	}
	user := &entity.User{
		Name:     userInput.Name,
		Email:    userInput.Email,
		Role:     entity.ROLE_DEFAULT,
		Password: userInput.Password,
	}
	userOutput, err := u.gateway.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return userOutput, nil
}
