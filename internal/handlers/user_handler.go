package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/mamun-jsx/user-management-golang/internal/models"
	"github.com/mamun-jsx/user-management-golang/internal/service"
)

type UserHandler struct {
	service *service.UserService
}

// make a new handler

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// create user handler

func (h *UserHandler) CreateUser(c fiber.Ctx) error {
	//* STEP : 1  declare a variable to store up-comming request data
	newUser := new(models.UserModel)

	//* STEP : 2 up-comming data land json formate. so Convert it into struct and set the sata into newUser variable
	if err := c.Bind().Body(newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	//* STEP : 3 call service to validate and match the logic we write
	if err := h.service.CreateUser(newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})

	}
	// send response to client
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
		"data":    newUser,
	})
}
