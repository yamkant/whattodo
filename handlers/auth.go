package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

type KakaoUserInfo struct {
	ID          int64     `json:"id"`
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

func GetLoginToken(c *gin.Context) {
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
	c.JSON(http.StatusOK, gin.H{"data": userInfo})
}
