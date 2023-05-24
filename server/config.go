package server

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Port                        string `json:"port"`
	AccessTokenSecret           string `json:"accessTokenSecret"`
	AccessTokenLifetimeMinutes  int    `json:"accessTokenLifetimeMinutes"`
	RefreshTokenSecret          string `json:"refreshTokenSecret"`
	RefreshTokenLifetimeMinutes int    `json:"refreshTokenLifetimeMinutes"`
}

func NewConfig() *Config {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	accessTokenLifetimeMinutes, err := strconv.Atoi(os.Getenv("ACCESS_TOKEN_LIFETIME_MINUTES"))
	if err != nil {
		log.Fatal("Error parsing ACCESS_TOKEN_LIFETIME_MINUTES:", err)
	}

	refreshTokenLifetimeMinutes, err := strconv.Atoi(os.Getenv("REFRESH_TOKEN_LIFETIME_MINUTES"))
	if err != nil {
		log.Fatal("Error parsing REFRESH_TOKEN_LIFETIME_MINUTES:", err)
	}

	return &Config{
		Port:                        os.Getenv("PORT"),
		AccessTokenSecret:           os.Getenv("ACCESS_TOKEN_SECRET"),
		AccessTokenLifetimeMinutes:  int(accessTokenLifetimeMinutes),
		RefreshTokenSecret:          os.Getenv("REFRESH_TOKEN_SECRET"),
		RefreshTokenLifetimeMinutes: int(refreshTokenLifetimeMinutes),
	}
}
