package repository

import (
	"database/sql"
	_ "embed"
	"fmt"

	"github.com/RensVanGiersbergen/skatetracker/models"
	log "github.com/sirupsen/logrus"
)

//go:embed queries/get_user_by_email.sql
var query string

// User queries
func GetUserByEmail(email string) (models.User, error) {
	// Create user object
	var user models.User

	// Execute the query
	err := db.QueryRow(query, email).Scan(&user.UserId, &user.Username, &user.Email, &user.HashedPassword, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Infof("No matches for email: %v", email)
			return models.User{}, fmt.Errorf("email not found: %w", sql.ErrNoRows)
		}
		log.Errorf("Error executing query: %v", err)
		return models.User{}, fmt.Errorf("internal error: %w", err)
	}
	return user, nil
}
