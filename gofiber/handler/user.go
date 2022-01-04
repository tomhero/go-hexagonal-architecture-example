package handler

import (
	"fmt"
	"gofiber/logs"
	"gofiber/service"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	userSrv service.UserService
}

func NewUserHandler(userSrv service.UserService) userHandler {
	return userHandler{userSrv: userSrv}
}

func (h userHandler) SignUpHandler(c *fiber.Ctx) error {
	userReq := service.UserRequest{}
	// data := map[string]interface{}{}
	err := c.BodyParser(&userReq)

	if err != nil {
		return err
	}
	logs.Info(fmt.Sprintf("data = %#v", userReq))
	respose, err := h.userSrv.Register(userReq)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(respose)
}

func (h userHandler) SignInHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("Login user Success")
}
