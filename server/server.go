package server

import (
	"log"
	"net/http"
)

func Start() {
	authHandler := NewAuthHandler()
	userHandler := NewUserHandler()

	http.HandleFunc("/login", authHandler.Login)
	http.HandleFunc("/profile", userHandler.GetProfile)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
