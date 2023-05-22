package server

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
)

type JwtCustomClaims struct {
	jwt.RegisteredClaims
	ID int `json:"id"`
}

func GenerateToken(userId, lifetimeMinutes int, secret string) (string, error) {
	claims := &JwtCustomClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(lifetimeMinutes))),
		},
		userId,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // important HS256

	return token.SignedString([]byte(secret))

}

func GetTokenFromBearerString(bearerString string) string {
	if bearerString == "" {
		return ""
	}

	parts := strings.Split(bearerString, "Bearer")
	if len(parts) != 2 {
		return ""
	}

	token := strings.TrimSpace(parts[1])
	if len(token) < 1 {
		return ""
	}
	return token
}

func ValidateToken(tokenstring, secret string) (*JwtCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenstring, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil // by default
	}) // code from documentation package

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JwtCustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New("failed to parse token claims")
	}

	return claims, nil
}
