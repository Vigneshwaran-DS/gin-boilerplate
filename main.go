package main

import (
	"flag"
	"fmt"
	"log"

	"gin-boilerplate/config"
	"gin-boilerplate/database"
	"gin-boilerplate/models"
	"gin-boilerplate/router"
	"gin-boilerplate/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	// 打印启动 banner
	utils.PrintBanner()

	// 解析命令行参数
	env := flag.String("e", "development", "运行环境 (development, production, test)")
	flag.Parse()

	// 加载配置
	config.LoadConfig(*env)

	// 设置 Gin 模式
	gin.SetMode(config.AppConfig.Server.Mode)

	// 初始化数据库
	database.InitDB()

	// 自动迁移数据库表
	if err := database.GetDB().AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// 设置路由
	r := router.SetupRouter()

	// 启动服务
	addr := fmt.Sprintf(":%s", config.AppConfig.Server.Port)
	log.Printf("Server is running on http://localhost%s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
