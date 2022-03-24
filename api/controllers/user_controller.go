package controllers

import (
	"net/http"

	"github.com/akselarzuman/go-jaeger/api/models"
	"github.com/akselarzuman/go-jaeger/internal/pkg/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserServiceInterface
}

func NewUserController() *UserController {
	return &UserController{
		userService: services.NewUserService(),
	}
}

func (ctr *UserController) Add(c *gin.Context) {
	var req models.CreateUserRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := ctr.userService.Add(c.Request.Context(), req.Name, req.Surname, req.Email, req.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
	})
}
