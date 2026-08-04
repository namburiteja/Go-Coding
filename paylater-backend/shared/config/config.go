package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// LoadEnv loads the first .env found among common locations.
func LoadEnv() {
	candidates := []string{".env"}

	if wd, err := os.Getwd(); err == nil {
		candidates = append(candidates,
			filepath.Join(wd, ".env"),
			filepath.Join(wd, "..", ".env"),
			filepath.Join(wd, "..", "..", ".env"),
		)
	}

	for _, path := range candidates {
		if err := godotenv.Load(path); err == nil {
			return
		}
	}

	log.Fatal("Error loading .env file")
}

func GetEnv(key string) string {
	return os.Getenv(key)
}
