package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	JWT      JWTConfig      `mapstructure:"jwt"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type JWTConfig struct {
	Secret     string `mapstructure:"secret"`
	ExpireTime int    `mapstructure:"expire_time"` // 单位：小时
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}

var AppConfig *Config

func LoadConfig(env string) {
	// 如果未指定环境，默认为 development
	if env == "" {
		env = "development"
	}

	// 加载默认配置
	defaultViper := viper.New()
	defaultViper.SetConfigName("default")
	defaultViper.SetConfigType("yaml")
	defaultViper.AddConfigPath("./config")
	defaultViper.AddConfigPath(".")

	if err := defaultViper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read default config file: %v", err)
	}
	log.Printf("Loaded default configuration from: %s", defaultViper.ConfigFileUsed())

	// 加载环境特定配置
	envViper := viper.New()
	envViper.SetConfigName(env)
	envViper.SetConfigType("yaml")
	envViper.AddConfigPath("./config")
	envViper.AddConfigPath(".")

	if err := envViper.ReadInConfig(); err != nil {
		log.Printf("No %s config file found, using default only", env)
	} else {
		log.Printf("Loaded %s configuration from: %s", env, envViper.ConfigFileUsed())
	}

	// 合并配置：默认配置 + 环境配置
	// 先解析默认配置
	AppConfig = &Config{}
	if err := defaultViper.Unmarshal(AppConfig); err != nil {
		log.Fatalf("Failed to unmarshal default config: %v", err)
	}

	// 合并环境配置（环境配置会覆盖默认配置）
	if err := envViper.Unmarshal(AppConfig); err != nil {
		log.Printf("Failed to unmarshal %s config: %v", env, err)
	}

	log.Printf("Configuration loaded successfully for environment: %s", env)
}
