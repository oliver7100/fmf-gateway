package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/oliver7100/advertisement-service/proto"
	"github.com/oliver7100/fmf-gateway/clients"
	uploadProto "github.com/oliver7100/upload-service/proto"
)

type advertisementController struct {
	advertisementClient proto.AdvertisementServiceClient
	uploadServiceClient uploadProto.UploadServiceClient
}

type DeleteAdvertisementRequestBody struct {
	Id int `json:"id"`
}

func (controller *advertisementController) createAdvertisement(c *fiber.Ctx) error {
	v := new(proto.Advertisement)

	file, _ := c.FormFile("file")

	clients.UploadFile(controller.uploadServiceClient, file)

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

	if err := c.QueryParser(v); err != nil {
		return fiber.NewError(500, err.Error())
	}

	res, err := controller.advertisementClient.GetAdvertisements(context.Background(), v)

	fmt.Println(res)

	if err != nil {
		return fiber.NewError(500, err.Error())
	}

	return c.JSON(res)
}

func (controller *advertisementController) DeleteAdvertisement(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)

	if err != nil {
		fiber.NewError(500, err.Error())
	}

	res, err := controller.advertisementClient.DisableAdvertisement(
		context.Background(),
		&proto.DisableAdvertisementRequest{
			Id: int32(id),
		},
	)

	if err != nil {
		return fiber.NewError(500, err.Error())
	}

	return c.JSON(res)
}

func newAdvertisementController(client proto.AdvertisementServiceClient, uploadServiceClient uploadProto.UploadServiceClient) *advertisementController {
	return &advertisementController{
		advertisementClient: client,
		uploadServiceClient: uploadServiceClient,
	}
}

func RegisterAdvertisementController(router fiber.Router, advertisementClient proto.AdvertisementServiceClient, uploadServiceClient uploadProto.UploadServiceClient) {
	advertisementRouter := router.Group("/advertisement")

	controller := newAdvertisementController(
		advertisementClient,
		uploadServiceClient,
	)

	advertisementRouter.Post("/", controller.createAdvertisement)
	advertisementRouter.Get("/", controller.FindAllAdvertisements)
	advertisementRouter.Delete("/:id<int>", controller.DeleteAdvertisement)
}
