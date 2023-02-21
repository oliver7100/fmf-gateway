package controllers

import (
	"context"
	"net/mail"

	"github.com/gofiber/fiber/v2"
	TokenProto "github.com/oliver7100/token-service/proto"
	UserProto "github.com/oliver7100/user-service/proto"
)

type authController struct {
	userClient  UserProto.UserServiceClient
	tokenClient TokenProto.AuthServiceClient
}

type LoginCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func validateMail(m string) bool {
	_, err := mail.ParseAddress(m)
	return err == nil
}

func (controller *authController) Register(c *fiber.Ctx) error {
	var s UserProto.CreateUserRequest

	if err := c.BodyParser(&s); err != nil {
		return fiber.NewError(500, "Post request invalid")
	}

	if ok := validateMail(s.User.Email); !ok {
		return fiber.NewError(500, "Invalid email")
	}

	r, err := controller.userClient.CreateUser(context.Background(), &s)

	if err != nil {
		return fiber.NewError(500, err.Error())
	}

	return c.JSON(r)
}

func (controller *authController) Login(c *fiber.Ctx) error {
	var creds LoginCredentials

	//var s UserProto.GetUserRequest

	if err := c.BodyParser(&creds); err != nil {
		return fiber.NewError(500, "Post request invalid")
	}

	_, err := controller.userClient.CanUserLogin(context.Background(), &UserProto.CanUserLoginRequest{
		Email:    creds.Email,
		Password: creds.Password,
	})

	if err != nil {
		return fiber.NewError(404, "Account not found.")
	}

	token, err := controller.tokenClient.GenerateToken(context.Background(), &TokenProto.GenerateTokenReqeust{
		Username: creds.Email,
	})

	if err != nil {
		return fiber.NewError(500, "Couldnt generate token")
	}

	return c.JSON(token)
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
	authRouter.Post("/login", controller.Login)
}
