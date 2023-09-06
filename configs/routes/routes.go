package routes

import (
	"example/hello/controller"

	"github.com/gofiber/fiber/v2"
)

func ToDoRoute(fiberClient *fiber.App) {

	route := fiberClient.Group("/task")

	route.Get("/get", controller.GetAllTask)
	route.Post("/create", controller.CreateTask)
	route.Put("/update/:id", controller.UpdateTask)

}

func OsuRoute(fiberClient *fiber.App) {

	route := fiberClient.Group("/osu")

	route.Post("/token", controller.GetNewToken)
	route.Get("/user/:userString", controller.GetUserDetail)
	route.Delete("/token", controller.DeleteExpiredToken)
	route.Get("/playCount/:userString",controller.DeadUserDataVisizulator)
}
