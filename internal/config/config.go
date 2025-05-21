package config

import (
	"os"
	"strconv"
)

type Config struct {
	UseMock    bool
	Port       string
	JWTSecret  string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
}

var AppConfig Config

// LoadConfig 從環境變量加載配置
func LoadConfig() {
	AppConfig = Config{
		UseMock:    getEnvAsBool("USE_MOCK", true), // 默認使用 mock 數據
		Port:       getEnvAsString("PORT", "9000"),
		JWTSecret:  getEnvAsString("JWT_SECRET", "your-secret-key"),
		DBHost:     getEnvAsString("DB_HOST", "localhost"),
		DBPort:     getEnvAsString("DB_PORT", "5432"),
		DBUser:     getEnvAsString("DB_USER", "postgres"),
		DBPassword: getEnvAsString("DB_PASSWORD", ""),
		DBName:     getEnvAsString("DB_NAME", "stray_map"),
		DBSSLMode:  getEnvAsString("DB_SSL_MODE", "disable"),
	}
}

func getEnvAsString(key string, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		boolValue, err := strconv.ParseBool(value)
		if err == nil {
			return boolValue
		}
	}
	return defaultValue
}
