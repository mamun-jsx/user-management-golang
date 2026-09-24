package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/mamun-jsx/user-management-golang/internal/app"
	"github.com/mamun-jsx/user-management-golang/internal/router"
)

func main() {

	application := app.BootstrapApp()
	fiberApp := fiber.New()
	router.ApplicationRouter(fiberApp, application)

	serverAddr := fmt.Sprintf("%s", application.Config.AppPort)
	log.Printf("🚀 Server is running on http://localhost%s", serverAddr)
	// handle error if have
	if err := fiberApp.Listen(serverAddr); err != nil {
		log.Fatalf("server failed to run: %v", err)
	}
}
