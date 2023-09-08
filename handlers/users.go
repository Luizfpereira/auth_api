package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsersHandler struct {
}

func NewUsersHandler() *UsersHandler {
	return &UsersHandler{}
}

func (u *UsersHandler) Register(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "test"})
}
