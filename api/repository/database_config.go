package repository

import (
	"database/sql"
	"embed"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

// PostgreSQL driver
var db *sql.DB

// embed queries from the queries directory into binary
//
//go:embed queries/*.sql
var queryFiles embed.FS

// QueryStore holds the content of embedded SQL files by name.
var queryStore = make(map[string]string)

// InitPostgresDB initializes the connection to a PostgreSQL database.
func InitPostgresDB(connStr string) {
	// Check if the connection string is provided
	if connStr == "" {
		log.Fatal("POSTGRES_CONNECTION_STRING environment variable is not set")
	}

	var err error
	for {
		// Open a connection to the PostgreSQL database
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			log.Errorf("Failed to open a DB connection: %v", err)
		} else {
			// Verify the connection is established successfully
			err = db.Ping()
			if err == nil {
				log.Info("Connected to PostgreSQL Instance")
				break
			} else {
				log.Errorf("Failed to connect to PostgreSQL: %v", err)
			}
		}

		log.Info("Retrying in 10 seconds...")
		time.Sleep(10 * time.Second)
	}
}

// ClosePostgresDB closes the connection to the PostgreSQL database
func ClosePostgresDB() {
	if db != nil {
		err := db.Close()
		if err != nil {
			log.Errorf("Error closing the DB connection: %v", err)
		} else {
			log.Info("PostgreSQL connection closed")
		}
	}
}

// LoadQueries loads all embedded queries into the QueryStore map.
func LoadQueries() error {
	entries, err := queryFiles.ReadDir("queries")
	if err != nil {
		return fmt.Errorf("failed to read query directory: %w", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			content, err := queryFiles.ReadFile("queries/" + entry.Name())
			if err != nil {
				return fmt.Errorf("failed to read query file %s: %w", entry.Name(), err)
			}
			queryStore[entry.Name()] = string(content)
		}
	}

	return nil
}
