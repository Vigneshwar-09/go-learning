package controller

import (
	"context"

	"example/hello/configs"
	"example/hello/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var taskCollection *mongo.Collection = configs.GetCollection(configs.DB, "task")

func GetAllTask(c *fiber.Ctx) error {
	var taskList [](model.Todo)

	cursor, err := taskCollection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "data": err})
	}

	if err = cursor.All(context.TODO(), &taskList); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "You are here", "data": taskList})
}

func CreateTask(c *fiber.Ctx) error {

	task := new(model.Todo)
	err := c.BodyParser(task)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	task.ID = uuid.New().String()
	task.Status = false

	taskCollection.InsertOne(c.Context(), task)
	return c.JSON(fiber.Map{"status": "success", "message": "Created task", "data": task})
}

func UpdateTask(c *fiber.Ctx) error {

	task := new(model.Todo)
	taskId := c.Params("id")
	// taskCollection.FindOne(c.Context(),bson.D{{Key: "_id",Value: taskId}}).Decode(&task)
	task.Status = true
	_, err := taskCollection.ReplaceOne(c.Context(), bson.M{"_id": taskId}, task)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Updated task", "data": task})
}
