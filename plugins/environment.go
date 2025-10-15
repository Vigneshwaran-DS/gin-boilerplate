package plugins

import (
	"gorm.io/gorm"
)

// PluginEnvironment plugin runtime environment, contains all dependencies required by plugins
type PluginEnvironment struct {
	// DB database connection
	DB *gorm.DB

	// Can extend more dependencies in the future:
	// Redis *redis.Client
}

// NewPluginEnvironment creates plugin environment
func NewPluginEnvironment(db *gorm.DB) *PluginEnvironment {
	return &PluginEnvironment{
		DB: db,
	}
}
