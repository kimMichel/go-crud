package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kimMichel/go-crud/controllers"
	"github.com/kimMichel/go-crud/database"
)

func init() {
	database.ConnectDB()
}

func main() {
	app := fiber.New()

	app.Post("/api/post", controllers.PostCreate)
	app.Get("/api/posts", controllers.PostsIndex)
	app.Get("/api/posts/:id", controllers.PostShow)

	app.Listen(":3000")
}
