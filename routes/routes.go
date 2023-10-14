package routes

import (
	"fmt"

	"example.com/m/middlewares"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// Server ... struct to hold global variables
type Server struct {
	Router *gin.Engine
}

func (server *Server) Init(port string) {
	gin.SetMode(gin.ReleaseMode)

	server.Router = gin.New()
	store := cookie.NewStore([]byte("secret"))
	server.Router.Use(sessions.Sessions("sessionid", store))
	server.Router.Use(middlewares.AuthMiddleware())

	server.initRoutes()
	server.authRoute()

	apiV1 := server.Router.Group("/api/v1")
	server.apiTodoRoute(apiV1)
	server.apiUserRoute(apiV1)
	server.apiStatusRoute(apiV1)

	server.Router.LoadHTMLGlob("views/*.html")
	server.Router.Static("/css", "views/css")
	server.Router.Static("/fonts", "views/fonts")
	server.Router.Static("/img", "views/img")
	server.Router.Static("/js", "views/js")

	fmt.Println("Server started.\nListening on 127.0.0.1/")
	server.Router.Run(port)
}
