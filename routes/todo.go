package routes

import "github.com/gin-gonic/gin"

func (server *Server) apiTodoRoute(routerGroup *gin.RouterGroup) {
	router := routerGroup.Group("/todos")
	router.GET("", GetTodos)
	router.POST("", AddTodo)
	router.PATCH("/:id/", UpdateTodo)
	router.DELETE("/:id/", DeleteTodo)
}
