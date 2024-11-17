package main

import (
	"os"

	"github.com/RensVanGiersbergen/skatetracker/router"
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
		// Set to info level if not in development
		log.SetLevel(log.InfoLevel)
	}

	// Router config
	r := router.SetupRouter()
	err := r.Run() // listen and serve on 0.0.0.0:8080 (Default)
	if err != nil {
		log.Fatal("Error starting server")
	}
}
