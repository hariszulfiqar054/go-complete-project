package utils

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string, defaultValue string) string {
	err := godotenv.Load(".env")
	if err != nil {
		return defaultValue
	}
	return os.Getenv(key)
}
