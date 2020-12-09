package config

import (
	"os"

	"github.com/joho/godotenv"
)

var env map[string]string

func init() {
	godotenv.Load()
}

// GetString gets an environment variable in string format.
func GetString(key string) string {
	return os.Getenv(key)
}
