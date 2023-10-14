package routes

import (
	"example.com/m/handlers"
	"github.com/gin-gonic/gin"
)

func authRoute(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/kakao", handlers.GetAuthCode)
	routerGroup.GET("/kakao/callback", handlers.GetAuthToken)

	routerGroup.GET("/login", handlers.Login)
}
