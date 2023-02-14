package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/oliver7100/user-service/proto"
)

type authController struct {
	userClient proto.UserServiceClient
}

func (controller *authController) Register(c *fiber.Ctx) error {
	var s proto.CreateUserRequest

	if err := c.BodyParser(&s); err != nil {
		return fiber.NewError(500, "Post request invalid")
	}

	fmt.Println(s)

	//controller.userClient.CreateUser(context.Background(), &s)

	return c.JSON(s)
}

func (controller *authController) Login(c *fiber.Ctx) error {
	var s proto.GetUserRequest

	if err := c.BodyParser(&s); err != nil {
		return fiber.NewError(500, "Post request invalid")
	}

	return c.JSON(s)
}

func newAuthController(userClient proto.UserServiceClient) *authController {
	return &authController{
		userClient,
	}
}

func RegisterAuthController(router fiber.Router, userClient proto.UserServiceClient) {
	authRouter := router.Group("/auth")

	controller := newAuthController(
		userClient,
	)

	authRouter.Post("/register", controller.Register)
}
