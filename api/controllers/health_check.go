package controllers

import (
	"log"
	"net/http"

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

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
