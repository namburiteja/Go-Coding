package config

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// LoadEnv loads the first existing meaningful file from paths into the process environment.
// OS environment variables already set (Docker/Kubernetes) are never overridden.
// If no file is found, loading is skipped so injected env still works.
//
// Callers should pass service/gateway .env candidates first when running from the
// repository root so a leftover root .env cannot win.
func LoadEnv(paths ...string) {
	if len(paths) == 0 {
		paths = []string{".env"}
	}

	for _, path := range paths {
		if !isMeaningfulEnvFile(path) {
			continue
		}
		if err := godotenv.Load(path); err == nil {
			log.Printf("config: loaded %s", path)
			return
		}
	}

	log.Printf("config: no .env file found among %v; using process environment", paths)
}

// isMeaningfulEnvFile reports whether path exists and contains at least one
// non-empty, non-comment line (so comment-only stubs do not win LoadEnv).
func isMeaningfulEnvFile(path string) bool {
	f, err := os.Open(path)
	if err != nil {
		return false
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		return true
	}
	return false
}

func GetEnv(key string) string {
	return os.Getenv(key)
}

// RequireEnv returns the value of key or exits if unset/empty.
func RequireEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("required environment variable %s is not set", key)
	}
	return value
}

// ListenAddr returns the HTTP listen address from PORT (e.g. "9091" or ":9091").
func ListenAddr() string {
	port := RequireEnv("PORT")
	if strings.HasPrefix(port, ":") {
		return port
	}
	return ":" + port
}
