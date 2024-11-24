package main

import (
	"os"

	"github.com/RensVanGiersbergen/skatetracker/repository"
	"github.com/RensVanGiersbergen/skatetracker/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Logging config
	// Force logging colors in terminal
	log.SetFormatter(&log.TextFormatter{
		ForceColors: true,
	})

	// Check if the --development flag is provided
	if len(os.Args) > 1 && os.Args[1] == "--development" {
		log.SetLevel(log.DebugLevel)
		err := godotenv.Load()
		log.Debug("Loading .env file")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	} else {
		// Set production mode
		gin.SetMode(gin.ReleaseMode)
		log.SetLevel(log.InfoLevel)
	}

	// Initialize the PostgreSQL connection
	repository.InitPostgresDB(os.Getenv("POSTGRES_CONNECTION_STRING"))

	defer repository.ClosePostgresDB() // Ensure the DB connection is closed when the app stops

	// Load all queries
	errQueries := repository.LoadQueries()
	if errQueries != nil {
		log.Fatalf("Error loading queries: %v", errQueries)
	}

	// Router config
	r := router.SetupRouter()
	errRouter := r.Run() // listen and serve on 0.0.0.0:8080 (Default)
	if errRouter != nil {
		log.Fatal("Error starting gin router")
	}
}
