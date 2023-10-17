package handlers

import (
	"time"

	"github.com/gin-gonic/gin"
)

func RenderHome(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"title": "What2Do",
	})
}

func RenderJoin(c *gin.Context) {
	c.HTML(200, "join.html", gin.H{
		"title": "What2Do",
	})
}

func Welcome(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Server started successfully at" + time.Now().String(),
	})
}
