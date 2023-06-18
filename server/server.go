package server

import (
	"golang-jwt/config"
	"golang-jwt/handler"
	"golang-jwt/middleware"
	"log"
	"net/http"
)

func Start(cfg *config.Config) {

	authHandler := handler.NewAuthHandler(cfg)
	userHandler := handler.NewUserHandler(cfg)
	refreshHandler := handler.NewRefreshHandler(cfg)

	http.HandleFunc("/login", authHandler.Login)
	http.HandleFunc("/profile", middleware.ValidateAccessToken(userHandler.GetProfile))
	http.HandleFunc("/register", authHandler.Register)
	http.HandleFunc("/refresh", middleware.ValidRefreshToken(refreshHandler.GetTokenPair))

	log.Fatal(http.ListenAndServe(cfg.Port, nil))
}
