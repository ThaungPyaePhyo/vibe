package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thaungpyaephyo/studentvibe/handlers"
	"github.com/thaungpyaephyo/studentvibe/middleware"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.POST("/register", handlers.CreateUser)
		api.POST("/login", handlers.LoginUser)
		api.POST("/logout", middleware.Auth() ,handlers.LogoutUser)
		api.GET("/posts", middleware.Auth(), handlers.GetPosts)
		api.POST("/posts", middleware.Auth(), handlers.CreatePost)
	}

	router.Static("/static", "./frontend/dist")
    router.NoRoute(func(c *gin.Context) {
        c.File("./frontend/dist/index.html")
    })
}