package main

import (
	"context"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/thaungpyaephyo/studentvibe/config"
	"github.com/thaungpyaephyo/studentvibe/handlers"
	"github.com/thaungpyaephyo/studentvibe/routes"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/gin-contrib/sessions"
    "github.com/gin-contrib/sessions/redis"
)

func main() {
    config.ConnectDB()
    router := gin.Default()
	secret := os.Getenv("SESSION_SECRET")
	store, err := redis.NewStore(10, "tcp", "localhost:6379", "", "", []byte(secret))
	if err != nil {
		panic(err)
	}
    router.Use(sessions.Sessions("mysession", store))
    routes.RegisterRoutes(router)
	client, _ := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
    handlers.UserCollection = client.Database("student_vibe").Collection("users")
    router.Run(":8000")
}