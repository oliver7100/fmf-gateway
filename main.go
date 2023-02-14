package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/oliver7100/fmf-gateway/clients"
	"github.com/oliver7100/fmf-gateway/controllers"
	"github.com/oliver7100/fmf-gateway/middleware"
)

func main() {
	app := fiber.New()

	api := app.Group("api")

	/* app.Use(
		jwtWare.New(
			jwtWare.Config{
				SigningKey: []byte("RandomString"),
			},
		),
	) */

	app.Use(middleware.New(middleware.Config{}))

	userServiceClient, err := clients.NewUserServiceClient(
		clients.NewConfig(
			":9000",
		),
	)

	authServiceClient, err := clients.NewTokenClient(
		clients.NewConfig(
			":8000",
		),
	)

	if err != nil {
		panic(err)
	}

	controllers.RegisterAuthController(
		api,
		userServiceClient,
		authServiceClient,
	)

	panic(app.Listen(":3000"))
}
