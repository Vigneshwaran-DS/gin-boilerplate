package guest

import (
	"gin-boilerplate/plugins"
	"gin-boilerplate/plugins/guest/controllers"
	"gin-boilerplate/plugins/guest/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Automatically register plugin during package initialization
func init() {
	plugins.Register("guest", NewGuestPlugin)
}

type GuestPlugin struct {
	db *gorm.DB
}

// NewGuestPlugin creates Guest plugin instance
func NewGuestPlugin(env *plugins.PluginEnvironment) plugins.Plugin {
	return &GuestPlugin{db: env.DB}
}

func (p *GuestPlugin) RouterPath() string {
	return "/guest"
}

func (p *GuestPlugin) Register(group *gin.RouterGroup) error {

	if err := p.db.AutoMigrate(&models.Guest{}); err != nil {
		return err
	}

	// Register routes
	controller := controllers.NewGuestController()
	group.POST("/login", controller.Login)

	return nil
}
