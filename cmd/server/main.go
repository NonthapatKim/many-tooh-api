package main

import (
	"fmt"
	"log"
	"os"

	"github.com/NonthapatKim/many-tooth-api/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Println("Warning: .env file not loaded")
		}
	}

	if err := server.RunServer(); err != nil {
		fmt.Println("Run server failed!", err)
		os.Exit(1)
	}

}
