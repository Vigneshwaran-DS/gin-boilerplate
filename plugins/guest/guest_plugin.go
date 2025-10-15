package guest

import (
	"log"
	"time"

	"gin-boilerplate/plugins"
	"gin-boilerplate/plugins/guest/controllers"
	"gin-boilerplate/plugins/guest/middleware"
	"gin-boilerplate/plugins/guest/models"
	"gin-boilerplate/plugins/guest/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Automatically register plugin during package initialization
func init() {
	plugins.Register("guest", NewGuestPlugin)
}

type GuestPlugin struct {
	db      *gorm.DB
	service *services.GuestService
}

// NewGuestPlugin creates Guest plugin instance
func NewGuestPlugin(env *plugins.PluginEnvironment) plugins.Plugin {
	service := services.NewGuestService(env.DB)
	return &GuestPlugin{
		db:      env.DB,
		service: service,
	}
}

func (p *GuestPlugin) RouterPath() string {
	return "/guest"
}

func (p *GuestPlugin) Register(group *gin.RouterGroup) error {
	// Auto-migrate guest table
	if err := p.db.AutoMigrate(&models.Guest{}); err != nil {
		return err
	}

	// Initialize controller
	controller := controllers.NewGuestController(p.service)

	// Public routes (no authentication required)
	group.POST("/login", controller.Login)

	// Protected routes (require guest JWT authentication)
	authenticated := group.Group("")
	authenticated.Use(middleware.GuestAuth())
	{
		authenticated.GET("/info", controller.GetCurrentGuest)
	}

	// Start cleanup task in background
	go p.startCleanupTask()

	log.Println("✅ Guest plugin registered successfully with cleanup task")

	return nil
}

// startCleanupTask runs periodic cleanup of inactive guest users
func (p *GuestPlugin) startCleanupTask() {
	ticker := time.NewTicker(24 * time.Hour) // Run every 24 hours
	defer ticker.Stop()

	// Run immediately on startup
	p.runCleanup()

	for range ticker.C {
		p.runCleanup()
	}
}

// runCleanup performs the actual cleanup operation
func (p *GuestPlugin) runCleanup() {
	// Remove guests inactive for more than 30 days
	inactiveDuration := 30 * 24 * time.Hour

	deleted, err := p.service.CleanupInactiveGuests(inactiveDuration)
	if err != nil {
		log.Printf("❌ Guest cleanup failed: %v", err)
		return
	}

	if deleted > 0 {
		log.Printf("🧹 Guest cleanup: removed %d inactive guests", deleted)
	}
}
