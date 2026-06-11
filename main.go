package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

// Config 配置结构体
type Config struct {
	Version   string
	Debug     bool
	OutputDir string
}

// DefaultConfig 默认配置
func DefaultConfig() *Config {
	return &Config{
		Version:   "1.0.0",
		Debug:     false,
		OutputDir: "./output",
	}
}

// Init 初始化
func (c *Config) Init() error {
	if err := os.MkdirAll(c.OutputDir, 0755); err != nil {
		return fmt.Errorf("创建目录失败: %w", err)
	}
	log.Printf("[INFO] 初始化完成 v%s", c.Version)
	return nil
}

func main() {
	config := DefaultConfig()
	if err := config.Init(); err != nil {
		log.Fatalf("[ERROR] %v", err)
	}
	log.Printf("[INFO] 启动时间: %s", time.Now().Format(time.RFC3339))
	fmt.Println("程序运行完成")
}
