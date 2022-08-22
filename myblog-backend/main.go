package main

import (
	"log"
	"os"

	"github.com/baiest/myblog/myblog-backend/models"
	"github.com/baiest/myblog/myblog-backend/server"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	dbConfig := models.DatabaseConfig{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Port:     os.Getenv("DB_PORT"),
		DbName:   os.Getenv("DB_NAME"),
	}

	serverConfig := models.ConfigServer{
		Port:           os.Getenv("PORT"),
		JwtKey:         os.Getenv("JWT_SECRET"),
		DatabaseConfig: dbConfig,
	}
	s := server.NewServer(serverConfig)
	s.Start()
}
