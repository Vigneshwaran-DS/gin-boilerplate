package router

import (
	"gin-boilerplate/controllers"
	"gin-boilerplate/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Use middleware
	r.Use(middleware.CORS())
	r.Use(middleware.Logger())

	// API version grouping
	v1 := r.Group("/api/v1")
	{
		// Health check
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status":  "ok",
				"message": "Service is running",
			})
		})

		// Authentication routes (no auth required)
		authController := controllers.NewAuthController()
		authRoutes := v1.Group("/auth")
		{
			authRoutes.POST("/register", authController.Register)
			authRoutes.POST("/login", authController.Login)
		}

		// Routes requiring authentication
		authenticated := v1.Group("")
		authenticated.Use(middleware.JWTAuth())
		{
			// Current user information
			authenticated.GET("/me", authController.GetCurrentUser)
			authenticated.PUT("/me", authController.UpdateCurrentUser)
		}
	}

	return r
}
