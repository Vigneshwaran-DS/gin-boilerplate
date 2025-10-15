package main

import (
	"flag"
	"fmt"
	"gin-boilerplate/plugins"
	"log"

	"gin-boilerplate/config"
	"gin-boilerplate/database"
	"gin-boilerplate/models"
	"gin-boilerplate/router"
	"gin-boilerplate/utils"

	_ "gin-boilerplate/plugins/guest" // Automatically register guest plugin

	"github.com/gin-gonic/gin"
)

func main() {
	// Print startup banner
	utils.PrintBanner()

	// Parse command line arguments
	env := flag.String("e", "development", "Runtime environment (development, production, test)")
	flag.Parse()

	// Load configuration
	config.LoadConfig(*env)

	// Set Gin mode
	gin.SetMode(config.AppConfig.Server.Mode)

	// Initialize database
	database.InitDB()

	// Auto migrate database tables
	if err := database.GetDB().AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Setup router
	r := router.SetupRouter()

	// Create plugin runtime environment
	pluginEnv := plugins.NewPluginEnvironment(database.GetDB())

	// Automatically load all registered plugins
	if err := plugins.LoadAllPlugins(r.Group("/api/v1/plugin"), pluginEnv); err != nil {
		log.Fatal("Failed to load plugins:", err)
	}

	// Start server
	addr := fmt.Sprintf(":%s", config.AppConfig.Server.Port)
	log.Printf("Server is running on http://localhost%s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
