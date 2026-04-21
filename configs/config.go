package configs

import "os"

type Config struct {
	JWTSecret string
}

func LoadConfig() *Config {
	return &Config{
		JWTSecret: getEnv("JWT_SECRET", "supersecret"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}