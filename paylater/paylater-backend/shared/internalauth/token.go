package internalauth

import "paylater/shared/config"

// Header is the HTTP header used for service-to-service authentication.
// JWT Authorization must never be used for internal calls.
const Header = "X-Internal-Service-Token"

// EnvKey is the environment variable holding the shared internal token.
const EnvKey = "INTERNAL_SERVICE_TOKEN"

// Token returns INTERNAL_SERVICE_TOKEN from the process environment.
func Token() string {
	return config.RequireEnv(EnvKey)
}
