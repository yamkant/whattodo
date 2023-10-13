package handlers

import (
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// RenderHome ... render the index.html page
func RenderHome(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"title": "Go Gin Boiler Plate",
	})
}

func Welcome(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Server started successfully at" + time.Now().String(),
	})
}

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

type TodoAddDTO struct {
	Content   string `json:"content"`
	Completed bool   `json:"completed"`
}

type TodoUpdateDTO struct {
	Completed bool      `json:"completed"`
	StartedAt time.Time `json:"started_at"`
	EndedAt   time.Time `json:"ended_at"`
}

type TodoDeleteDTO struct {
	Content   string `json:"content"`
	Completed bool   `json:"completed"`
}

func GetTodos(c *gin.Context) {
	sort.Slice(todos, func(i, j int) bool {
		return todos[i].ID > todos[j].ID
	})
	c.IndentedJSON(http.StatusOK, todos)
}

func AddTodo(c *gin.Context) {
	bodyData := TodoAddDTO{}
	if err := c.ShouldBind(&bodyData); err != nil {
		return
	}

	id := len(todos) + 1
	todo := &Todo{id, bodyData.Content, bodyData.Completed, time.Time{}, time.Time{}, time.Now()}
	todos = append(todos, *todo)
	c.IndentedJSON(http.StatusCreated, todo)
}

func UpdateTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	bodyData := TodoUpdateDTO{}
	if err := c.ShouldBind(&bodyData); err != nil {
		return
	}

	for i := range todos {
		if todos[i].ID == id {
			todos[i].Completed = bodyData.Completed
			todos[i].StartedAt = bodyData.StartedAt
			todos[i].EndedAt = bodyData.EndedAt
			c.IndentedJSON(http.StatusOK, todos[i])
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}

func DeleteTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	bodyData := TodoDeleteDTO{}
	if err := c.ShouldBind(&bodyData); err != nil {
		return
	}

	for i := range todos {
		if todos[i].ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			c.IndentedJSON(http.StatusOK, nil)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}
