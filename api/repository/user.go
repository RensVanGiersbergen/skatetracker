package repository

import (
	"database/sql"
	_ "embed"
	"errors"
	"fmt"

	"github.com/RensVanGiersbergen/skatetracker/models"
	"github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

// User queries
func GetUserByEmail(email string) (models.User, error) {
	var user models.User

	// Execute the query
	err := db.QueryRow(queryStore["get_user_by_email.sql"], email).Scan(&user.UserId, &user.Username, &user.Email, &user.HashedPassword, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Debugf("No matches for email: %v", email)
			return models.User{}, fmt.Errorf("email not found: %w", sql.ErrNoRows)
		}
		log.Errorf("Error executing query: %v", err)
		return models.User{}, fmt.Errorf("internal error: %w", err)
	}
	return user, nil
}

func AddUser(register models.RegisterRequest) error {
	// Execute the query
	_, err := db.Exec(queryStore["add_user.sql"], register.Username, register.Email, register.Password)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			switch pqErr.Code {
			case "23505": // Unique violation in PostgreSQL
				if pqErr.Constraint == "users_username_key" { // Adjust for your actual constraint name
					return fmt.Errorf("username already exists")
				}
				if pqErr.Constraint == "users_email_key" { // Adjust for your actual constraint name
					return fmt.Errorf("email already exists")
				}
			}
		}
		return fmt.Errorf("error executing query: %w", err)
	}
	return nil
}
