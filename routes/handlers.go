package routes

import (
	"net/http"
	"time"

	"example.com/m/models"
	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	var todos []models.Todo
	models.DB.Order("id desc").Find(&todos)

	c.JSON(http.StatusOK, gin.H{"data": todos})
}

func AddTodo(c *gin.Context) {
	var bodyData models.TodoAddDTO
	if err := c.ShouldBindJSON(&bodyData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := models.Todo{Content: bodyData.Content, Completed: bodyData.Completed, StartedAt: time.Time{}, EndedAt: time.Time{}, CreatedAt: time.Now()}
	models.DB.Create(&todo)

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func UpdateTodo(c *gin.Context) {
	var todo models.Todo

	if err := models.DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var bodyData models.TodoUpdateDTO
	if err := c.ShouldBindJSON(&bodyData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	models.DB.Model(&todo).Updates(bodyData)
	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func DeleteTodo(c *gin.Context) {
	var todo models.Todo

	if err := models.DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"data": true})
}

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
