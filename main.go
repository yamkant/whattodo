package main

import (
	"example.com/m/models"
	"example.com/m/routes"
)

var server = routes.Server{}

func main() {
	models.ConnectDatabase()
	server.Init(":8080")
}
