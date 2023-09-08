package usecase

import "time"

type UserUsecase struct {
}

type UserInput struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
}

type UserOutput struct {
	ID        int64     `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Role      string    `json:"role,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUserUsecase() *UserUsecase {
	return &UserUsecase{}
}

func (u *UserUsecase) CreateUser(userInput UserInput) {
	
}
