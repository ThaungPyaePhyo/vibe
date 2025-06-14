package handlers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"github.com/thaungpyaephyo/studentvibe/models"
	"go.mongodb.org/mongo-driver/mongo"
)

var LikeCollection *mongo.Collection

func LikePost(c *gin.Context) {
	postID := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID format"})
		return
	}
	session := sessions.Default(c)
	userID := session.Get("user_id")
	userIDStr, ok := userID.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user_id in session"})
		return
	}
	userObjID, err := primitive.ObjectIDFromHex(userIDStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user_id format"})
		return
	}
	like := models.Like{
		ID:        primitive.NewObjectID(),
		PostID:    objID,
		UserID:    userObjID,
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),		
	}

	_, err = LikeCollection.InsertOne(c.Request.Context(), like)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like post"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Post liked successfully", "like": like})

}