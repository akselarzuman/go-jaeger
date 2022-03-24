package controllers

import (
	"net/http"

	"github.com/akselarzuman/go-jaeger/internal/pkg/opentelemetry"
	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	span := opentelemetry.SpanFromContext(c.Request.Context())
	defer span.End()

	opentelemetry.AddSpanEvents(span, "test_logs", map[string]string{
		"name":    "Aksel",
		"surname": "Arzuman",
	})

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
