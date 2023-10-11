package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	Completed bool      `json:"completed"`
	StartedAt time.Time `json:"started_at"`
	EndedAt   time.Time `json:"ended_at"`
	CreatedAt time.Time `json:"created_at"`
}

var todos = []Todo{
	{ID: 1, Content: "First todo", Completed: true, StartedAt: time.Now(), EndedAt: time.Now(), CreatedAt: time.Now()},
	{ID: 2, Content: "Second todo", Completed: true, StartedAt: time.Now(), EndedAt: time.Now(), CreatedAt: time.Now()},
	{ID: 3, Content: "Third todo", Completed: false, StartedAt: time.Now(), EndedAt: time.Now(), CreatedAt: time.Now()},
}

type TodoDTO struct {
	Content   string `json:"content"`
	Completed bool   `json:"completed"`
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func addTodo(c *gin.Context) {
	bodyData := TodoDTO{}
	if err := c.ShouldBind(&bodyData); err != nil {
		return
	}

	id := len(todos) + 1
	todo := &Todo{id, bodyData.Content, bodyData.Completed, time.Time{}, time.Time{}, time.Now()}
	todos = append(todos, *todo)
	c.IndentedJSON(http.StatusCreated, todo)
}

func updateTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	bodyData := TodoDTO{}
	if err := c.ShouldBind(&bodyData); err != nil {
		return
	}

	for i := range todos {
		if todos[i].ID == id {
			todos[i].Content = bodyData.Content
			todos[i].Completed = bodyData.Completed
			c.IndentedJSON(http.StatusOK, todos[i])
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "todo.html", gin.H{
			"title": "What to do?",
		})
	})
	router.GET("/api/v1/todos/", getTodos)
	router.POST("/api/v1/todos/", addTodo)
	router.PATCH("/api/v1/todos/:id/", updateTodo)
	router.Run(":8080")
}
