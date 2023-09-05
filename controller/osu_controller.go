package controller

import (
	"example/hello/service"

	"github.com/gofiber/fiber/v2"
)

func GetNewToken(c *fiber.Ctx) error {

	authToken, err := service.FetchAuthToken()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err, "discription": "Shit broke"})
	}
	return c.Status(200).JSON(fiber.Map{"data": authToken})
}
