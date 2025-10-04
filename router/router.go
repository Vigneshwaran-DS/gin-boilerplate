package router

import (
	"gin-boilerplate/controllers"
	"gin-boilerplate/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 使用中间件
	r.Use(middleware.CORS())
	r.Use(middleware.Logger())

	// API 版本分组
	v1 := r.Group("/api/v1")
	{
		// 健康检查
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status":  "ok",
				"message": "Service is running",
			})
		})

		// 认证路由（无需鉴权）
		authController := controllers.NewAuthController()
		authRoutes := v1.Group("/auth")
		{
			authRoutes.POST("/register", authController.Register)
			authRoutes.POST("/login", authController.Login)
		}

		// 需要认证的路由
		authenticated := v1.Group("")
		authenticated.Use(middleware.JWTAuth())
		{
			// 当前用户信息
			authenticated.GET("/me", authController.GetCurrentUser)
			authenticated.PUT("/me", authController.UpdateCurrentUser)
		}
	}

	return r
}
