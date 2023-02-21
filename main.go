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
			":7000",
		),
	)

	controllers.RegisterAuthController(
		api,
		userServiceClient,
		authServiceClient,
	)

	controllers.RegisterAdvertisementController(
		api,
		advertisementClient,
	)

	panic(app.Listen(":3000"))
}
