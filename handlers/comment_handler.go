package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/thaungpyaephyo/studentvibe/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
	"github.com/gin-contrib/sessions"
	"go.mongodb.org/mongo-driver/mongo"
)

var CommentCollection *mongo.Collection

func CommentPost(c *gin.Context) {
	postID := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID format"})
		return
	}
	session := sessions.Default(c)
	userID := session.Get("user_id")
	if userID == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userIDStr, ok := userID.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user_id in session"})
		return
	}

	userObjID, err := primitive.ObjectIDFromHex(userIDStr)
	var comment models.Comment
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user_id format"})
		return
	}
	comment.ID = primitive.NewObjectID()
	comment.PostID = objID
	comment.UserID = userObjID
	comment.CreatedAt = time.Now().Format(time.RFC3339)
	comment.UpdatedAt = comment.CreatedAt
	comment.Content = c.PostForm("content")
	_, err = CommentCollection.InsertOne(c.Request.Context(), comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to comment on post"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Comment added successfully", "comment": comment})	

	
}