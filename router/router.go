package router

import (
	"auth_api/handlers"
	"auth_api/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitializeRouter(middleware *middleware.Middleware, usersHandler *handlers.UsersHandler) *gin.Engine {
	router := gin.Default()
	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{"message": "page not found"}) })
	router.GET("/", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, "up and running...") })

	auth := router.Group("/auth")
	auth.POST("/register", usersHandler.Register)
	auth.POST("/login", usersHandler.SignInUser)
	auth.GET("/logout", usersHandler.Logout)
	auth.GET("/refresh", usersHandler.Refresh)

	router.GET("/users/me", middleware.Authorize, usersHandler.GetMe)

	return router
}
