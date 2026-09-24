package router

import (
	"github.com/gofiber/fiber/v3"
	"github.com/mamun-jsx/user-management-golang/internal/app"
)




func ApplicationRouter(fiberApp *fiber.App, application *app.App) {
	fiberApp.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
