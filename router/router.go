package router

import (
	"github.com/gin-gonic/gin"
	"github.com/akashkrao99/go-sample-http/internal/health"
)

func GetRouter()(*gin.Engine) {

	router := gin.Default()

	router.GET("/health", health.GetHealth)

	return router
}
