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
	response, err := h.userSrv.Register(userReq)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}

func (h userHandler) SignInHandler(c *fiber.Ctx) error {
	userReq := new(service.UserLoginRequest)
	// data := map[string]interface{}{}
	err := c.BodyParser(&userReq)

	if err != nil {
		return err
	}

	response, err := h.userSrv.Login(userReq.Username, userReq.Password)
	if err != nil {
		return err
	}
	response.Username = userReq.Username

	return c.Status(fiber.StatusOK).JSON(response)
}
