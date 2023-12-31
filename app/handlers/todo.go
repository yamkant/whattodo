package handlers

import (
	"fmt"
	"net/http"
	"time"

	"example.com/m/models"
	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	var todos []models.Todo
	models.DB.Raw(`
		SELECT
			*
		FROM todos
		WHERE user_id = 1 AND deleted_at IS NULL
		ORDER BY
		CASE
			WHEN completed = true THEN 1 AND 2
			ELSE 2
		END, completed_at DESC, created_at DESC
	`, user.ID).Scan(&todos)

	c.JSON(http.StatusOK, gin.H{"data": todos})
}

func AddTodo(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	var bodyData models.TodoAddDTO
	if err := c.ShouldBindJSON(&bodyData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := models.Todo{
		Content:   bodyData.Content,
		Completed: bodyData.Completed,
		UserID:    user.ID,
	}
	models.DB.Create(&todo)

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func UpdateTodo(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	var todo models.Todo
	if err := models.DB.Where("id = ?", c.Param("id")).Where("user_id = ?", user.ID).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var bodyData models.TodoUpdateDTO
	if err := c.ShouldBindJSON(&bodyData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if bodyData.Completed {
		bodyData.CompletedAt = time.Now()
	} else {
		bodyData.CompletedAt = time.Time{}
	}
	models.DB.Model(&todo).Select("*").Updates(bodyData)

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func UpdateTodoContent(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	var todo models.Todo
	if err := models.DB.Where("id = ?", c.Param("id")).Where("user_id = ?", user.ID).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var bodyData models.TodoContentUpdateDTO
	if err := c.ShouldBindJSON(&bodyData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	models.DB.Model(&todo).Select("*").Updates(bodyData)

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func UpdateTodoStartAt(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	var todo models.Todo
	if err := models.DB.Where("id = ?", c.Param("id")).Where("user_id = ?", user.ID).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var bodyData models.TodoStartAtUpdateDTO
	bodyData.StartedAt = time.Now()
	models.DB.Model(&todo).Select("*").Updates(bodyData)

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func UpdateTodoEndAt(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	var todo models.Todo
	if err := models.DB.Where("id = ?", c.Param("id")).Where("user_id = ?", user.ID).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var bodyData models.TodoEndAtUpdateDTO
	if (todo.StartedAt != time.Time{}) {
		bodyData.Completed = true
		bodyData.EndedAt = time.Now()
		bodyData.CompletedAt = time.Now()
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "started_at is necessary for changing ended_at"})
	}
	models.DB.Model(&todo).Select("*").Updates(bodyData)

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func DeleteTodo(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	if err := models.DB.Where("id = ?", c.Param("id")).Where("user_id = ?", user.ID).Delete(&models.Todo{}).Error; err != nil {
		fmt.Println("제거 실패")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	// models.DB.Delete(&todo)
	fmt.Println("제거 성공")

	c.JSON(http.StatusOK, gin.H{"data": true})
}
