package main

import (
	"fmt"
	"golang-jwt/config"
	"golang-jwt/server"
)

func main() {
	cfg := config.NewConfig()
	fmt.Printf("server is runnig")
	server.Start(cfg)
}
