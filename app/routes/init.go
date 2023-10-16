package routes

import (
	"example.com/m/handlers"
	"github.com/gin-gonic/gin"
)

func initRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/", handlers.RenderHome)
	routerGroup.GET("/auth/join", handlers.RenderJoin)
}
