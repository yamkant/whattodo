package routes

import "github.com/gin-gonic/gin"

func (server *Server) initRoutes() {
	server.Router.GET("/", RenderHome)
}

func (server *Server) apiStatusRoute(routerGroup *gin.RouterGroup) {
	routerGroup.GET("", Welcome)
}
