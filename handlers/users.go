package handlers

import (
	"auth_api/auth"
	"auth_api/entity"
	"auth_api/errorConst"
	"auth_api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsersHandler struct {
	Usecase *usecase.UserUsecase
}

func NewUsersHandler(usecase *usecase.UserUsecase) *UsersHandler {
	return &UsersHandler{Usecase: usecase}
}

func (u *UsersHandler) Register(ctx *gin.Context) {
	name := ctx.PostForm("name")
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")
	passwordConfirm := ctx.PostForm("password_confirm")

	if name == "" || email == "" || password == "" || passwordConfirm == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "incomplete form"})
		return
	}
	if password != passwordConfirm {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "password and confirmation password don't match"})
		return
	}
	if !entity.IsPasswordValid(password) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errorConst.PASSWORD_INVALID})
		return
	}

	hashedPassword := auth.HashPassword(password)

	input := usecase.UserInput{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	}
	output, err := u.Usecase.CreateUser(input)
	if err != nil {
		if err.Error() == errorConst.EMAIL_REGISTERED {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, output)
}
