package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thaungpyaephyo/studentvibe/config"
)

func main() {
	config.ConnectDB()
	router :=  gin.Default()
	router.Static("/static", "./frontend/dist")
	router.NoRoute(func(c *gin.Context) {
		c.File("./frontend/dist/index.html")
	})
	api := router.Group("/api")
	{
		api.GET("/hello", func(c *gin.Context) {
			c.String(200, "Hello, World!")
		})
	}
	router.Run(":8080")
}