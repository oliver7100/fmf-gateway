package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/oliver7100/fmf-gateway/clients"
	"github.com/oliver7100/fmf-gateway/controllers"
)

func main() {
	app := fiber.New()

	api := app.Group("api")

	userServiceClient, _ := clients.NewUserServiceClient(
		clients.NewConfig(
			":9000",
		),
	)

	authServiceClient, _ := clients.NewTokenClient(
		clients.NewConfig(
			":8000",
		),
	)

	advertisementClient, _ := clients.NewAdvertisementClient(
		clients.NewConfig(
			":1338",
		),
	)

	uploadServiceClient, _ := clients.NewUploadServiceClient(
		clients.NewConfig(
			":1337",
		),
	)

	controllers.RegisterAuthController(
		api,
		userServiceClient,
		authServiceClient,
	)

	/* app.Use(
		jwtware.New(
			jwtware.Config{
				SigningKey: []byte("secret"),
			},
		),
	) */

	controllers.RegisterAdvertisementController(
		api,
		advertisementClient,
		uploadServiceClient,
	)

	panic(app.Listen(":3000"))
}
