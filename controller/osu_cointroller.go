package controller

import (
	"example/hello/service"

	"github.com/gofiber/fiber/v2"
)

func GetNewToken(c *fiber.Ctx) error {

	authToken := service.FetchAuthToken()

	return c.Status(200).JSON(fiber.Map{"data": authToken})
}
