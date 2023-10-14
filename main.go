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

	fmt.Println("Server started.\nListening on 127.0.0.1/")
	router.Run(":8080")
}
