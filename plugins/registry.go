package plugins

import (
	"log"
	"sync"

	"github.com/gin-gonic/gin"
)

// PluginFactory plugin factory function type, used to create plugin instances
type PluginFactory func(env *PluginEnvironment) Plugin

// pluginRegistry plugin registry
type pluginRegistry struct {
	mu      sync.RWMutex
	plugins map[string]PluginFactory
	order   []string // Maintain plugin registration order
}

// Global registry instance
var registry = &pluginRegistry{
	plugins: make(map[string]PluginFactory),
	order:   make([]string, 0),
}

// Register registers plugin to global registry
// This function should be called in the plugin's init() function
func Register(name string, factory PluginFactory) {
	registry.mu.Lock()
	defer registry.mu.Unlock()

	if _, exists := registry.plugins[name]; exists {
		log.Printf("⚠️  Plugin [%s] already registered, skipping...", name)
		return
	}

	registry.plugins[name] = factory
	registry.order = append(registry.order, name)
	log.Printf("📝 Plugin [%s] registered", name)
}

// LoadAllPlugins loads all registered plugins
func LoadAllPlugins(baseRouter *gin.RouterGroup, env *PluginEnvironment) error {
	registry.mu.RLock()
	defer registry.mu.RUnlock()

	log.Printf("🧩 Starting to load %d plugins...", len(registry.plugins))

	// Load plugins in registration order
	for _, name := range registry.order {
		factory := registry.plugins[name]

		// Create plugin instance
		plugin := factory(env)

		// Create plugin router group
		pluginRouter := baseRouter.Group(plugin.RouterPath())

		// Register plugin routes
		if err := plugin.Register(pluginRouter); err != nil {
			log.Printf("❌ Plugin [%s] failed to load: %v", name, err)
			return err
		}

		log.Printf("✅ Plugin [%s] loaded successfully", name)
	}

	log.Printf("🎉 All plugins loaded successfully!")
	return nil
}

// GetRegisteredPlugins gets all registered plugin names (for debugging)
func GetRegisteredPlugins() []string {
	registry.mu.RLock()
	defer registry.mu.RUnlock()

	result := make([]string, len(registry.order))
	copy(result, registry.order)
	return result
}
