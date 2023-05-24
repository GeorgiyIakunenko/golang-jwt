package server

import (
	"log"
	"net/http"
)

func Start() {

	cfg := NewConfig()

	authHandler := NewAuthHandler(cfg)
	userHandler := NewUserHandler(cfg)

	http.HandleFunc("/login", authHandler.Login)
	http.HandleFunc("/profile", userHandler.GetProfile)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
