package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"time"

	"example.com/m/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

type KakaoUserInfo struct {
	ID          uint64    `json:"id"`
	ConnectedAt time.Time `json:"connected_at"`
}

var (
	kakaoOauthConfig = oauth2.Config{
		ClientID:     os.Getenv("KAKAO_CLIENT_ID"),
		ClientSecret: os.Getenv("KAKAO_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("KAKAO_REDIRECT_URL"),
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://kauth.kakao.com/oauth/authorize",
			TokenURL: "https://kauth.kakao.com/oauth/token",
		},
	}
)

func GetAuthCode(c *gin.Context) {
	authURL := kakaoOauthConfig.AuthCodeURL("", oauth2.AccessTypeOffline)
	c.Redirect(http.StatusFound, authURL)
}

func GetAuthToken(c *gin.Context) {
	code := c.DefaultQuery("code", "")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "OAuth code is missing"})
		return
	}

	token, err := kakaoOauthConfig.Exchange(c, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange code for token"})
		return
	}

	session := sessions.Default(c)
	session.Set("access_token", token.AccessToken)
	session.Save()
	c.Redirect(http.StatusFound, "/auth/login")
}

func Login(c *gin.Context) {
	session := sessions.Default(c)
	accessToken := session.Get("access_token")
	if accessToken == nil {
		c.Redirect(http.StatusFound, "/auth/kakao")
	}

	userInfoURL := "https://kapi.kakao.com/v2/user/me"
	req, _ := http.NewRequest("GET", userInfoURL, nil)
	req.Header.Set("Authorization", "Bearer "+accessToken.(string))

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch kakao user info"})
		return
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to io read for kakao user info response"})
		return
	}

	var userInfo KakaoUserInfo
	err = json.Unmarshal(responseBody, &userInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal user info to object"})
		return
	}
	session.Set("kakao_user_id", userInfo.ID)
	session.Save()

	var user models.User
	err = models.DB.Where("kakao_id = ?", userInfo.ID).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.Redirect(http.StatusTemporaryRedirect, "/auth/join")
		return
	}
	c.Redirect(http.StatusTemporaryRedirect, "/")
}
