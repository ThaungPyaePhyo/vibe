package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thaungpyaephyo/studentvibe/config"
)

func main() {
	config.ConnectDB()
	router :=  gin.Default()
	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})
	router.Run(":8080")
}