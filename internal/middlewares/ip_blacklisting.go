package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IPBlacklistMiddleware(blacklist []string) gin.HandlerFunc {
    return func(c *gin.Context) {
        clientIP := c.ClientIP() // Get the client's IP address

        // Check if the client's IP is in the blacklist
        for _, ip := range blacklist {
            if clientIP == ip {
                c.JSON(http.StatusForbidden, gin.H{"error": "Your IP is blacklisted."})
                c.Abort() // Abort the request
                return
            }
        }

        c.Next() // Proceed to the next middleware/handler
    }
}