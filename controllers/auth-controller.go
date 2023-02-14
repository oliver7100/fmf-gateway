package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	TokenProto "github.com/oliver7100/token-service/proto"
	UserProto "github.com/oliver7100/user-service/proto"
)

type authController struct {
	userClient  UserProto.UserServiceClient
	tokenClient TokenProto.AuthServiceClient
}

func (controller *authController) Register(c *fiber.Ctx) error {
	var s UserProto.CreateUserRequest

	if err := c.BodyParser(&s); err != nil {
		return fiber.NewError(500, "Post request invalid")
	}

	fmt.Println(s)

	//controller.userClient.CreateUser(context.Background(), &s)

	return c.JSON(s)
}

func (controller *authController) Login(c *fiber.Ctx) error {
	var s UserProto.GetUserRequest

	if err := c.BodyParser(&s); err != nil {
		return fiber.NewError(500, "Post request invalid")
	}

	return c.JSON(s)
}

func newAuthController(userClient UserProto.UserServiceClient, tokenClient TokenProto.AuthServiceClient) *authController {
	return &authController{
		userClient,
		tokenClient,
	}
}

func RegisterAuthController(router fiber.Router, userClient UserProto.UserServiceClient, authClient TokenProto.AuthServiceClient) {
	authRouter := router.Group("/auth")

	controller := newAuthController(
		userClient,
		authClient,
	)

	authRouter.Post("/register", controller.Register)
}
