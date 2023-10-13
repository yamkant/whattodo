package main

import (
	"example.com/m/handlers"
	"example.com/m/models"
)

var server = handlers.Server{}

func main() {
	models.ConnectDatabase()
	server.Init(":8080")
}
