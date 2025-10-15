package controllers

import (
	"gin-boilerplate/plugins/guest/services"

	"github.com/gin-gonic/gin"
)

type GuestController struct {
	service *services.GuestService
}

func NewGuestController() *GuestController {
	return &GuestController{service: &services.GuestService{}}
}

func (c *GuestController) Login(ctx *gin.Context) {
	// Handle guest login logic
	ctx.JSON(200, gin.H{"message": "Guest login successful"})
}
