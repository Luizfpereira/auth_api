package middleware

import (
	"auth_api/auth"
	"auth_api/config"
	"auth_api/usecase"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Middleware struct {
	config      config.Config
	userUsecase *usecase.UserUsecase
}

func NewMiddleware(config config.Config, userUsecase *usecase.UserUsecase) *Middleware {
	return &Middleware{config: config, userUsecase: userUsecase}
}

func (m *Middleware) Authorize(ctx *gin.Context) {
	var accessToken string
	authorizationHeader := ctx.Request.Header.Get("Authorization")
	fields := strings.Fields(authorizationHeader)

	if len(fields) != 0 && fields[0] == "Bearer" {
		accessToken = fields[1]
	}

	if accessToken == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "user unauthorized"})
		return
	}

	sub, err := auth.ValidateToken(accessToken, m.config.AccessTokenPublicKey)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "invalid token"})
		return
	}

	id, err := strconv.Atoi(fmt.Sprint(sub))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "user unauthorized"})
		return
	}

	user, err := m.userUsecase.GetUserById(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "The user belonging to this token no logger exists"})
		return
	}

	ctx.Set("currentUser", user)
	ctx.Next()
}
