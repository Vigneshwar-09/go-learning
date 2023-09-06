package main

import (
	"example/hello/configs"
	"example/hello/configs/routes"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Hello, World!")

	app := fiber.New(configs.FiberConfig())
	routes.ToDoRoute(app)
	routes.OsuRoute(app)

	configs.ConnectDB()

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})
	app.Listen(":6000")

}
