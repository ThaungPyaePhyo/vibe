package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/thaungpyaephyo/studentvibe/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var UserCollection *mongo.Collection
var validate = validator.New()

func CreateUser(c *gin.Context) {
	var user  models.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return 
	}
	//validation
	if err := validate.Struct(user); err != nil {
		errs := err.(validator.ValidationErrors)
		errorMap := make(map[string]string)
		for _, e := range errs {
			errorMap[e.Field()] = e.Tag()
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Validation failed",
			"errors":  errorMap,
		})
		return
	} 

	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now().Format(time.RFC3339)
	user.UpdatedAt = user.CreatedAt
	
	_, err := UserCollection.InsertOne(context.TODO(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	} 
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": user})

}