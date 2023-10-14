package middlewares

import (
	"net/http"
	"strings"

	"example.com/m/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var whiteList = []string{
	"/api/v1/users",
}

func isWhiteList(url string) bool {
	if strings.HasPrefix(url, "/auth") {
		return true
	}
	for _, v := range whiteList {
		if v == url {
			return true
		}
	}
	return false
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if isWhiteList(c.Request.URL.Path) {
			c.Next()
		}

		session := sessions.Default(c)

		kakaoUserID := session.Get("kakao_user_id")
		if kakaoUserID == nil {
			c.Redirect(http.StatusFound, "/auth/kakao")
			c.Abort()
			return
		}

		var user models.User
		err := models.DB.Where("kakao_id = ?", kakaoUserID).First(&user).Error
		if err != nil {
			c.Redirect(http.StatusFound, "/auth/kakao")
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
