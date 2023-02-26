package controllers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	UserProto "github.com/oliver7100/user-service/proto"
)

type UserController struct {
	UserClient UserProto.UserServiceClient
}

func (c *UserController) GetUser(ctx *fiber.Ctx) error {
	if res, err := c.UserClient.GetUser(context.Background(), &UserProto.GetUserRequest{
		Identifier: &UserProto.GetUserRequest_Username{
			Username: "test",
		},
	}); err != nil {
		return fiber.NewError(404, err.Error())
	} else {
		return ctx.JSON(res)
	}
}
