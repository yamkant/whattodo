package main

import (
	"fmt"

	"example.com/m/models"
	"example.com/m/routes"
)

var server = routes.Server{}

func main() {
	models.ConnectDatabase()
	router := server.Init()

	fmt.Println("Server started.\n")
	router.Run("0.0.0.0:8080")
}
