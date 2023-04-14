package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kimMichel/go-crud/database"
)

func init() {
	database.ConnectDB()
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"message": "Hello, World!"})
	})

	app.Listen(":3000")
}
