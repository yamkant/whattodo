package routes

import (
	"example.com/m/handlers"
	"github.com/gin-gonic/gin"
)

func apiTodoRoute(routerGroup *gin.RouterGroup) {
	router := routerGroup.Group("/todos")
	router.GET("", handlers.GetTodos)
	router.POST("", handlers.AddTodo)
	router.PATCH("/:id/", handlers.UpdateTodo)
	router.PATCH("/:id/content/", handlers.UpdateContentTodo)
	router.PATCH("/:id/time/", handlers.UpdateTimeTodo)
	router.DELETE("/:id/", handlers.DeleteTodo)
}
