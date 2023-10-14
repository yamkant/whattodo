package routes

import (
	"example.com/m/handlers"
	"github.com/gin-gonic/gin"
)

func (server *Server) initRoutes() {
	server.Router.GET("/", handlers.RenderHome)
}

func (server *Server) apiStatusRoute(routerGroup *gin.RouterGroup) {
	routerGroup.GET("", handlers.Welcome)
}
