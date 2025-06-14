package middleware

import (
    "net/http"
    "github.com/gin-contrib/sessions"
    "github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
    return func(c *gin.Context) {
        session := sessions.Default(c)
        userID := session.Get("user_id")
        if userID == nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }
        sessionToken, _ := session.Get("xsrf_token").(string)
        headerToken := c.GetHeader("X-XSRF-TOKEN")
        if sessionToken == "" || headerToken == "" || sessionToken != headerToken {
            c.JSON(http.StatusForbidden, gin.H{"error": "Invalid or missing XSRF token"})
            c.Abort()
            return
        }
        c.Next()
    }
}