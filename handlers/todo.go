package handlers

import (
	"net/http"
	"time"

	"example.com/m/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	var todos []models.Todo
	// USER VALIDATION PROCESS
	session := sessions.Default(c)
	kakaoUserID := session.Get("kakao_user_id")
	if kakaoUserID == nil {
		c.Redirect(http.StatusFound, "/auth/kakao")
	}
	var tmpUser models.User
	err := models.DB.Where("kakao_id = ?", kakaoUserID).First(&tmpUser).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": tmpUser})
		return
	}

	models.DB.Where("user_id = ?", tmpUser.ID).Order("id desc").Find(&todos)

	c.JSON(http.StatusOK, gin.H{"data": todos})
}

func AddTodo(c *gin.Context) {
	var bodyData models.TodoAddDTO
	if err := c.ShouldBindJSON(&bodyData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// USER VALIDATION PROCESS
	session := sessions.Default(c)
	kakaoUserID := session.Get("kakao_user_id")
	if kakaoUserID == nil {
		c.Redirect(http.StatusFound, "/auth/kakao")
	}
	var tmpUser models.User
	err := models.DB.Where("kakao_id = ?", kakaoUserID).First(&tmpUser).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": tmpUser})
		return
	}

	todo := models.Todo{Content: bodyData.Content, Completed: bodyData.Completed, StartedAt: time.Time{}, EndedAt: time.Time{}, CreatedAt: time.Now(), UserID: tmpUser.ID}
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
