package router

import (
	"os"

	"github.com/akselarzuman/go-jaeger/api/controllers"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func Setup() *gin.Engine {
	app := gin.New()
	app.Use(otelgin.Middleware(os.Getenv("JAEGER_SERVICE_NAME")))
	app.GET("/health_check", controllers.HealthCheck)

	return app
}
