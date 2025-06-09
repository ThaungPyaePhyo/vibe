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
)

func main() {
    config.ConnectDB()
    router := gin.Default()
    routes.RegisterRoutes(router)
	client, _ := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
    handlers.UserCollection = client.Database("student_vibe").Collection("users")
    router.Run(":8000")
}