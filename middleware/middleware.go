package middleware

import (
	"context"
	"golang-jwt/config"
	"golang-jwt/server/response"
	"golang-jwt/server/service"
	"net/http"
)

func ValidateAccessToken(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		AuthHeader := r.Header.Get("Authorization")
		token := service.GetTokenFromBearerString(AuthHeader)

		claims, err := service.ValidateToken(token, config.NewConfig().AccessSecret)
		if err != nil {
			response.SendInvalidCredentials(w)
			return
		}

		ctx := context.WithValue(r.Context(), config.NewConfig().AccessSecret, claims)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}

func ValidRefreshToken(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		refreshTokenString := r.Header.Get("Authorization")

		claims, err := service.ValidateToken(refreshTokenString, config.NewConfig().RefreshSecret)
		if err != nil {
			response.SendInvalidCredentials(w)
			return
		}

		c := context.WithValue(r.Context(), config.NewConfig().RefreshSecret, claims)
		req := r.WithContext(c)

		next.ServeHTTP(w, req)
	})
}
