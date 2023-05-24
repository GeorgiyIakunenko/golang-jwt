package server

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

const (
	//port                        = ":8080"
	accessTokenSecret           = "access_token_secret"
	refreshTokenSecret          = "refresh_token_secret"
	accessTokenLifetimeMinutes  = 3
	refreshTokenLifetimeMinutes = 180
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

	return &Config{
		Port:                        os.Getenv("PORT"),
		AccessTokenSecret:           accessTokenSecret,
		AccessTokenLifetimeMinutes:  accessTokenLifetimeMinutes,
		RefreshTokenSecret:          refreshTokenSecret,
		RefreshTokenLifetimeMinutes: refreshTokenLifetimeMinutes,
	}
}
