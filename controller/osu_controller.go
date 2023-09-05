package controller

import (
	"example/hello/service"

	"github.com/gofiber/fiber/v2"
)

func GetNewToken(c *fiber.Ctx) error {

	authToken, err := service.FetchAuthToken()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err, "description": "Shit broke"})
	}
	return c.Status(200).JSON(fiber.Map{"data": authToken})
}

func GetUserDetail(c *fiber.Ctx) error {
	response, err := service.FetchUserData(c.Params("userString"))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err, "description": "Shit broke"})
	}
	return c.Status(200).JSON(fiber.Map{"data": response})
}

func DeleteExpiredToken(c *fiber.Ctx) error {
	err := service.RemoveExpiredToken()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err, "description": "Shit broke"})
	}
	return c.Status(200).JSON(fiber.Map{"status": "sucess"})
}
