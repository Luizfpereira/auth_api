package usecase

import (
	"auth_api/entity"
	"auth_api/errorConst"
	"auth_api/gateway"
	"errors"
	"strings"
)

type UserUsecase struct {
	gateway gateway.UserGateway
}

type UserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUserUsecase(gateway gateway.UserGateway) *UserUsecase {
	return &UserUsecase{gateway: gateway}
}

func (u *UserUsecase) CreateUser(userInput UserInput) (*entity.UserOutput, error) {
	user := &entity.User{
		Name:     userInput.Name,
		Email:    userInput.Email,
		Role:     entity.ROLE_DEFAULT,
		Password: userInput.Password,
	}
	userOutput, err := u.gateway.CreateUser(user)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return nil, errors.New(errorConst.EMAIL_REGISTERED)
		}
		return nil, err
	}
	return userOutput, nil
}

func (u *UserUsecase) GetUserByEmail(email string) (*entity.User, error) {
	user, err := u.gateway.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
