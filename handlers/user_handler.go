package handlers

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gin-contrib/sessions"
	"github.com/thaungpyaephyo/studentvibe/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/v2/bson"
	"golang.org/x/crypto/bcrypt"
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

	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
        return
    }

	user.Password = string(hashed)
	
	_, err = UserCollection.InsertOne(context.TODO(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	} 
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": user})
}

func LoginUser(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return 
	} 

	var user models.User
	err := UserCollection.FindOne(context.TODO(), bson.M{"email": req.Email}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email "})
		return
	} 

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
			"bcrypt_error": err.Error(),
		})
		return
	}
	session := sessions.Default(c)
	session.Set("user_id", user.ID.Hex())

	csrfToken, err := generateCSRFToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate CSRF token"})
		return
	}
	session.Set("xsrf_token", csrfToken)
	session.Save()
	c.SetCookie("XSRF-TOKEN", csrfToken, 3600, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": user})
}

func generateCSRFToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err 
	}
	return base64.StdEncoding.EncodeToString(b), nil
}