package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

// PrintBanner prints startup banner
func PrintBanner() {
	bannerPath := filepath.Join("config", "banner.txt")

	// Read banner file
	content, err := os.ReadFile(bannerPath)
	if err != nil {
		// If file doesn't exist or read fails, use default banner
		fmt.Println("=== Gin Boilerplate ===")
		return
	}

	// Print banner
	fmt.Println(string(content))
}
