package main

import (
	"example.com/m/handlers"
)

var server = handlers.Server{}

func main() {
	// router := gin.Default()
	// router.LoadHTMLGlob("templates/**/*")
	// router.LoadHTMLGlob("views/*.html")
	// router.Static("/js", "views/js")
	// router.StaticFS("/static", http.Dir("static"))
	// router.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "todo.html", gin.H{
	// 		"title": "What to do?",
	// 	})
	// })

	server.Init(":8080")
}
