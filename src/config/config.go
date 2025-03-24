package config

import (
	"os"
	"strconv"
)

type Config struct {
	Addr string `json:"addr,omitempty"`
	Port string `json:"port,omitempty"`
}

func LoadConfig() Config {
	return Config{
		Port: getEnv("SERVER_PORT", "8080"),
		Addr: getEnv("SERVER_ADDR", ""),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvAsBool(key string, defaultValue bool) bool {
	valueStr := getEnv(key, "")
	if valueStr == "" {
		return defaultValue
	}

	value, err := strconv.ParseBool(valueStr)
	if err != nil {
		return defaultValue
	}

	return value
}
