package main

import (
	"github.com/jm61/restApi/initializers"
	"github.com/jm61/restApi/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
