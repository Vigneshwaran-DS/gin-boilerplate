package plugins

import (
	"github.com/gin-gonic/gin"
)

// Plugin plugin interface definition
type Plugin interface {
	// Register registers plugin routes and functionality
	Register(group *gin.RouterGroup) error

	// RouterPath returns the plugin's route prefix
	RouterPath() string
}
