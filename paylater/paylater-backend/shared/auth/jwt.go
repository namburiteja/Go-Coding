package auth

import (
	"time"

	"paylater/shared/config"

	"github.com/golang-jwt/jwt/v5"
)

const (
	RoleAdmin    = "ADMIN"
	RoleMerchant = "MERCHANT"
	RoleCustomer = "CUSTOMER"
)

type Claims struct {
	UserID int32  `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID int32, role string) (string, error) {
	secret := []byte(config.RequireEnv("JWT_SECRET"))
	expiry := config.RequireEnv("JWT_EXPIRY")

	duration, err := time.ParseDuration(expiry)
	if err != nil {
		return "", err
	}

	claims := Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func ValidateJWT(tokenString string) (*Claims, error) {
	secret := []byte(config.RequireEnv("JWT_SECRET"))

	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, jwt.ErrTokenInvalidClaims
	}

	return claims, nil
}
