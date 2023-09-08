package main

import (
	"auth_api/config"
	"auth_api/database"
	"auth_api/handlers"
	"auth_api/server"
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

	usersHandler := handlers.NewUsersHandler()
	server := server.NewServer(":" + config.Port)

	server.AddHandler("/", "GET", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, "up and running...") })
	server.Router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{"message": "page not found"}) })
	server.AddHandler("/users/register", "POST", usersHandler.Register)
	server.Start()
}
