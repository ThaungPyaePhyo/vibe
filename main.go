package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thaungpyaephyo/studentvibe/config"
)

func main() {
    config.ConnectDB()
    router := gin.Default()

    // API routes
    api := router.Group("/api")
    {
        api.GET("/hello", func(c *gin.Context) {
            c.String(200, "Hello, World!")
        })
    }

    // Serve static files under /static
    router.Static("/static", "./frontend/dist")

    // SPA fallback: serve index.html for all other routes
    router.NoRoute(func(c *gin.Context) {
        c.File("./frontend/dist/index.html")
    })

    router.Run(":8080")
}