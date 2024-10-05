package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetCorsMiddleware() gin.HandlerFunc {

	config := cors.Config{
		AllowOrigins:     []string{"http://example.com"}, //Replace with actual values
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 3600,
	}

	return cors.New(config)
}
