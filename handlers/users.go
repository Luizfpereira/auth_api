package handlers

import (
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
	u.Usecase.CreateUser()
	ctx.JSON(http.StatusOK, gin.H{"message": "test"})
}
