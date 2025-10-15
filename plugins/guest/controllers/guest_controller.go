package controllers

import (
	"net/http"

	"gin-boilerplate/plugins/guest/services"
	"gin-boilerplate/utils"
	"github.com/gin-gonic/gin"
)

type GuestController struct {
	service *services.GuestService
}

func NewGuestController(service *services.GuestService) *GuestController {
	return &GuestController{service: service}
}

// GuestLoginRequest represents the request body for guest login
type GuestLoginRequest struct {
	DeviceID   string `json:"device_id"`
	DeviceInfo string `json:"device_info"`
}

// Login handles guest auto-registration and login
func (c *GuestController) Login(ctx *gin.Context) {
	var req GuestLoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request data")
		return
	}

	// Get client IP address
	ipAddress := ctx.ClientIP()

	// Create guest user
	guest, token, err := c.service.CreateGuest(req.DeviceID, req.DeviceInfo, ipAddress)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to create guest user")
		return
	}

	utils.SuccessResponse(ctx, gin.H{
		"token":     token, 
		"guest_uid": guest.GuestUID,
		"guest": gin.H{
			"id":              guest.ID,
			"guest_uid":       guest.GuestUID,
			"device_id":       guest.DeviceID,
			"last_active_at":  guest.LastActiveAt,
			"created_at":      guest.CreatedAt,
		},
	})
}

// GetCurrentGuest returns the current guest information
func (c *GuestController) GetCurrentGuest(ctx *gin.Context) {
	// Get guest UID from context (set by guest auth middleware)
	guestUID, exists := ctx.Get("guest_uid")
	if !exists {
		utils.ErrorResponse(ctx, http.StatusUnauthorized, "Guest not authenticated")
		return
	}

	guestUIDStr := guestUID.(string)

	// Get guest information
	guest, err := c.service.GetGuestByUID(guestUIDStr)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}

	// Update last active timestamp
	c.service.UpdateLastActive(guestUIDStr)

	utils.SuccessResponse(ctx, gin.H{
		"id":             guest.ID,
		"guest_uid":      guest.GuestUID,
		"device_id":      guest.DeviceID,
		"device_info":    guest.DeviceInfo,
		"last_active_at": guest.LastActiveAt,
		"created_at":     guest.CreatedAt,
	})
}
