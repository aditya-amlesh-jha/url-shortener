package config

import (
	"os"
)

type Config struct {
	ServerPort string
	MySQLURI   string
	RedisAddr  string
}

func getEnv(key string, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func LoadConfig() *Config {
	return &Config{
		ServerPort: getEnv("SERVER_PORT", ":8080"),
		MySQLURI:   getEnv("NYSQLURI", ""),
		RedisAddr:  getEnv("RedisAddr", "localhost:6379"),
	}
}
