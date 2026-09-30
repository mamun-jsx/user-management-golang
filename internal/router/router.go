package router

import (
	"github.com/gofiber/fiber/v3"
	"github.com/mamun-jsx/user-management-golang/internal/app"
	"github.com/mamun-jsx/user-management-golang/internal/handlers"
	"github.com/mamun-jsx/user-management-golang/internal/repository"
	"github.com/mamun-jsx/user-management-golang/internal/service"
)

func ApplicationRouter(fiberApp *fiber.App, application *app.App) {
	fiberApp.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World User Management app !")
	})

	// Initialize dependencies
	userRepo := repository.NewUserRepository(application.DB)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	api := fiberApp.Group("/api")
	// user api endpoints
	api.Post("/create-user", userHandler.CreateUser)
}
