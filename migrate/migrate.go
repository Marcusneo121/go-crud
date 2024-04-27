package main

import (
	"github.com/marcus121neo/go-crud/initializers"
	"github.com/marcus121neo/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}