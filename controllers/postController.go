package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kimMichel/go-crud/database"
	"github.com/kimMichel/go-crud/models"
)

func PostCreate(c *fiber.Ctx) error {
	// Get data off req body
	var body struct {
		Title string
		Body  string
	}

	c.BodyParser(&body)

	// Create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := database.DB.Create(&post)

	if result.Error != nil {
		log.Fatal(c.Status(400))
	}

	// Return it
	return c.Status(200).JSON(fiber.Map{
		"post": post,
	})
}

func PostsIndex(c *fiber.Ctx) error {
	// Get the posts
	var posts []models.Post
	database.DB.Find(&posts)

	// Return with them
	return c.Status(200).JSON(fiber.Map{
		"post": posts,
	})
}

func PostShow(c *fiber.Ctx) error {
	// Get id off url
	id := c.Params("id")

	// Get the post
	var post models.Post
	database.DB.First(&post, id)

	// Respond with them
	return c.Status(200).JSON(fiber.Map{
		"post": post,
	})
}

func PostUpdate(c *fiber.Ctx) error {
	// Get id off url
	id := c.Params("id")

	// Get the data from request body
	var body struct {
		Title string
		Body  string
	}

	c.BodyParser(&body)

	// Find Post were updating
	var post models.Post
	database.DB.First(&post, id)

	// Update it
	database.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// Respond with them
	return c.Status(200).JSON(fiber.Map{
		"post": post,
	})
}

func PostDelete(c *fiber.Ctx) error {
	// Get the id off the url
	id := c.Params("id")

	// Delete the post
	database.DB.Delete(&models.Post{}, id)

	// Respond
	return c.SendStatus(200)
}
