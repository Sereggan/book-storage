package config

import (
	"os"
)

type Config struct {
	ServerAddress string
	GrpcPort      string
	DbAddress     string
}

func New() *Config {
	return &Config{
		ServerAddress: getEnv("SERVER_ADDRESS", "localhost:8080"),
		GrpcPort:      getEnv("GRPC_SERVER_PORT", "5300"),
		DbAddress:     getEnv("DB_ADDRESS", "host=localhost dbname=book_storage sslmode=disable user=sergey password=password"),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
