package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Server ... struct to hold global variables
type Server struct {
	Router *gin.Engine
}

func (server *Server) Init(port string) {
	gin.SetMode(gin.ReleaseMode)
	server.Router = gin.New()

	server.initRoutes()
	server.apiStatus()
	server.apiTodoRoutes()

	fmt.Println("Static setting")
	server.Router.LoadHTMLGlob("views/*.html")
	server.Router.Static("/css", "views/css")
	server.Router.Static("/fonts", "views/fonts")
	server.Router.Static("/img", "views/img")
	server.Router.Static("/js", "views/js")

	fmt.Println("Server started.\nListening on 127.0.0.1/")
	server.Router.Run(port)
}
