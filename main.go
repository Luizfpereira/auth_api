package main

import (
	"auth_api/config"
	"auth_api/database"
	"auth_api/handlers"
	"auth_api/middleware"
	"auth_api/repository"
	"auth_api/router"
	"auth_api/server"
	"auth_api/usecase"
	"log"
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

	middleware := middleware.NewMiddleware(config, userUsecase)

	router := router.InitializeRouter(middleware, usersHandler)

	server := server.NewServer(":"+config.Port, router)

	server.Start()
}
