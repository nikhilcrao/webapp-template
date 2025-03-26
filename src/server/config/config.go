package config

import (
	"os"
)

type Config struct {
	Addr string `json:"addr"`
	Port string `json:"port"`

	DatabasePath string `json:"database_path"`

	OAuthCallbackURL   string `json:"oauth_callback_domain"`
	GoogleClientID     string `json:"google_client_id"`
	GoogleClientSecret string `json:"google_client_secret"`
}

func LoadConfig() Config {
	return Config{
		Port:               getEnv("SERVER_PORT", "8080"),
		Addr:               getEnv("SERVER_ADDR", ""),
		DatabasePath:       getEnv("DATABASE_PATH", "postgres://localhost:5432/postgres"),
		OAuthCallbackURL:   getEnv("OAUTH_CALLBACK_DOMAIN", "http://localhost:8080/api/auth/google/callback"),
		GoogleClientID:     getEnv("GOOGLE_CLIENT_ID", ""),
		GoogleClientSecret: getEnv("GOOGLE_CLIENT_SECRET", ""),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

/*
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
*/
