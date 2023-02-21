package controllers

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/oliver7100/advertisement-service/proto"
)

type advertisementController struct {
	advertisementClient proto.AdvertisementServiceClient
}

func (controller *advertisementController) createAdvertisement(c *fiber.Ctx) error {
	v := new(proto.Advertisement)

	if err := c.BodyParser(v); err != nil {
		return fiber.NewError(500, err.Error())
	}

	res, err := controller.advertisementClient.CreateAdvertisement(context.Background(), v)

	if err != nil {
		return fiber.NewError(500, err.Error())
	}

	return c.JSON(res)
}

func (controller *advertisementController) FindAllAdvertisements(c *fiber.Ctx) error {
	v := new(proto.GetAllAdvertisementsRequest)

	if err := c.BodyParser(v); err != nil {
		return fiber.NewError(500, err.Error())
	}

	res, err := controller.advertisementClient.GetAdvertisements(context.Background(), v)

	fmt.Println(res)

	if err != nil {
		return fiber.NewError(500, err.Error())
	}

	return c.JSON(res)
}

func newAdvertisementController(client proto.AdvertisementServiceClient) *advertisementController {
	return &advertisementController{
		advertisementClient: client,
	}
}

func RegisterAdvertisementController(router fiber.Router, client proto.AdvertisementServiceClient) {
	advertisementRouter := router.Group("/advertisement")

	controller := newAdvertisementController(
		client,
	)

	advertisementRouter.Post("/", controller.createAdvertisement)
	advertisementRouter.Get("/", controller.FindAllAdvertisements)
}
