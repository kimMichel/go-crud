package main

import (
	"github.com/kimMichel/go-crud/database"
	"github.com/kimMichel/go-crud/models"
)

func init() {
	database.ConnectDB()
}

func main() {
	database.DB.AutoMigrate(&models.Post{})
}
