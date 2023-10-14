package middlewares

import (
	"net/http"

	"example.com/m/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		kakaoUserID := session.Get("kakao_user_id")
		if kakaoUserID == nil {
			c.Redirect(http.StatusFound, "/auth/kakao")
		}

		var user models.User
		err := models.DB.Where("kakao_id = ?", kakaoUserID).First(&user).Error
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
