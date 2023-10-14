package handlers

import (
	"net/http"
	"time"

	"example.com/m/models"
	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	var todos []models.Todo

	user := c.MustGet("user").(models.User)

	models.DB.Where("user_id = ?", user.ID).Order("id desc").Find(&todos)

	c.JSON(http.StatusOK, gin.H{"data": todos})
}

func AddTodo(c *gin.Context) {
	var bodyData models.TodoAddDTO
	if err := c.ShouldBindJSON(&bodyData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := c.MustGet("user").(models.User)

	todo := models.Todo{
		Content:   bodyData.Content,
		Completed: bodyData.Completed,
		StartedAt: time.Time{},
		EndedAt:   time.Time{},
		CreatedAt: time.Now(),
		UserID:    user.ID,
	}
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
