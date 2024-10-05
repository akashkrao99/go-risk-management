package router

import (
	"github.com/akashkrao99/go-sample-http/internal/health"
	"github.com/akashkrao99/go-sample-http/internal/middlewares"
	"github.com/akashkrao99/go-sample-http/internal/risks"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	
	router := gin.Default()

	// middlewares
	router.Use(middlewares.GetCorsMiddleware())

	// routes
	router.GET("/health", health.GetHealth)

	v1 := router.Group("/v1")
	{
		setupRisksRoutes(v1.Group("/risks"))
	}

	return router
}

func setupRisksRoutes(risksRouterGroup *gin.RouterGroup) {
	risksController := risks.NewRisksController()

	risksRouterGroup.OPTIONS("/")
	risksRouterGroup.GET("/", risksController.GetRisks)

	risksRouterGroup.OPTIONS("/:id")
	risksRouterGroup.GET("/:id", risksController.GetRiskById)

	risksRouterGroup.POST("/", risksController.CreateRisk)
}
