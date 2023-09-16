package main

import (
	"auth_api/config"
	"auth_api/database"
	"auth_api/handlers"
	"auth_api/repository"
	"auth_api/server"
	"auth_api/usecase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	instance := database.ConnectSingleton(config.PostgresConn)
	database.Migrate(instance)

	userRepo := repository.NewUserRepositoryPSQL(instance)

	userUsecase := usecase.NewUserUsecase(userRepo)

	usersHandler := handlers.NewUsersHandler(userUsecase, config)

	server := server.NewServer(":" + config.Port)
	server.AddHandler("/", "GET", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, "up and running...") })
	server.Router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{"message": "page not found"}) })
	server.AddHandler("/auth/register", "POST", usersHandler.Register)
	server.AddHandler("/auth/login", "POST", usersHandler.SignInUser)
	server.Start()
}
