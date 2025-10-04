package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

// PrintBanner 打印启动 banner
func PrintBanner() {
	bannerPath := filepath.Join("config", "banner.txt")

	// 读取 banner 文件
	content, err := os.ReadFile(bannerPath)
	if err != nil {
		// 如果文件不存在或读取失败，使用默认 banner
		fmt.Println("=== Gin Boilerplate ===")
		return
	}

	// 打印 banner
	fmt.Println(string(content))
}
