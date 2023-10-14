package routes

import (
	"example.com/m/handlers"
	"github.com/gin-gonic/gin"
)

func (server *Server) apiUserRoute(routerGroup *gin.RouterGroup) {
	router := routerGroup.Group("/users")
	router.POST("", handlers.AddUser)
}
