package api

import (
	"log"
	"os"

	"github.com/azinudinachzab/bq-loan-backend/api/controllers"
	"github.com/joho/godotenv"
)

func Run() {

	if os.Getenv("ENV") == "DEV" {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Error getting env, %v", err)
			return
		}
	}

	var server = controllers.Server{}
	if err := server.Initialize(
		controllers.DBConnection{
			Driver:   os.Getenv("DB_DRIVER"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Port:     os.Getenv("DB_PORT"),
			Host:     os.Getenv("DB_HOST"),
			Name:     os.Getenv("DB_NAME"),
		},
	); err != nil {
		log.Fatalf("Cannot connect to %v database", err)
		return
	}

	// seed.Load(server.DB)
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}
	log.Printf("Server started on port %s", port)
	server.Run(port)
}
