package routes

import (
	"example.com/m/handlers"
)

func (server *Server) authRoute() {
	router := server.Router.Group("/auth")

	router.GET("/kakao", handlers.GetAuthCode)
	router.GET("/kakao/callback", handlers.GetAuthToken)

	router.GET("/login", handlers.GetLoginToken)
}
