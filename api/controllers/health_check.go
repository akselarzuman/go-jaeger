package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	jlog "github.com/opentracing/opentracing-go/log"
)

func HealthCheck(c *gin.Context) {
	span, err := opentracing.StartSpanFromContext(c.Request.Context(), "health_check_controller")
	if err != nil {
		log.Println(err.Err())
	}
	defer span.Finish()

	span.LogFields(
		jlog.String("name", "Aksel"),
		jlog.String("surname", "Arzuman"),
	)

	c.JSON(200, gin.H{
		"message": "ok",
	})
}
