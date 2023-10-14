package routes

import (
	"example.com/m/middlewares"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Router *gin.Engine
}

func setUpRoute() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("sessionid", store))
	router.Use(middlewares.AuthMiddleware())

	initV1 := router.Group("/")
	initRoutes(initV1)

	authV1 := router.Group("/auth")
	authRoute(authV1)

	apiV1 := router.Group("/api/v1")
	apiTodoRoute(apiV1)
	apiUserRoute(apiV1)
	return router
}

func setUpStatic(router *gin.Engine) {
	router.LoadHTMLGlob("views/*.html")
	router.Static("/css", "views/css")
	router.Static("/fonts", "views/fonts")
	router.Static("/img", "views/img")
	router.Static("/js", "views/js")
}

func (server *Server) Init() *gin.Engine {
	server.Router = setUpRoute()
	setUpStatic(server.Router)

	return server.Router
}
