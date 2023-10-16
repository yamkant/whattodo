package handlers

import (
	"net/http"

	"example.com/m/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context) {
	session := sessions.Default(c)
	kakaoUserID := session.Get("kakao_user_id")
	if kakaoUserID == nil {
		c.Redirect(http.StatusFound, "/auth/kakao")
	}

	var tmpUser models.User
	err := models.DB.Where("kakao_id = ?", kakaoUserID).First(&tmpUser).Error
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": tmpUser})
	}

	user := models.User{KakaoID: kakaoUserID.(uint64)}
	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}
