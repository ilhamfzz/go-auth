package main

import (
	"authentication/internal/api"
	"authentication/internal/component"
	"authentication/internal/config"
	"authentication/internal/middleware"
	"authentication/internal/repository"
	"authentication/internal/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cnf := config.Get()
	dbConnection := component.DatabaseConnection(cnf)
	cacheConnection := component.GetCacheConnection()
	userRepository := repository.NewUser(dbConnection)
	userService := service.NewUser(userRepository, cacheConnection)
	authMiddleware := middleware.Authenticate(userService)

	app := fiber.New()
	api.NewAuth(app, userService, authMiddleware)

	_ = app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
}
