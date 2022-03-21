package router

import "github.com/gin-gonic/gin"

func Setup() *gin.Engine {
	app := gin.New()

	return app
}
