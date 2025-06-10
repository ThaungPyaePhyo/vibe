package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thaungpyaephyo/studentvibe/handlers"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.POST("/register", handlers.CreateUser)
		api.POST("/login", handlers.LoginUser)
	}
	router.Static("/static", "./frontend/dist")

    router.NoRoute(func(c *gin.Context) {
        c.File("./frontend/dist/index.html")
    })
}