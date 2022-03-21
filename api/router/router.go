package router

import (
	"github.com/akselarzuman/go-jaeger/api/controllers"
	"github.com/akselarzuman/go-jaeger/api/middlewares"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	app := gin.New()
	app.Use(middlewares.OpentracingMiddlewareGin())
	app.GET("/health_check", controllers.HealthCheck)

	return app
}
