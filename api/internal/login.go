package internal

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/RensVanGiersbergen/skatetracker/models"
	"github.com/RensVanGiersbergen/skatetracker/repository"
	"golang.org/x/crypto/bcrypt"

	log "github.com/sirupsen/logrus"
)

var (
	// Errors
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

func Login(loginRequest models.LoginRequest) (string, error) {
	// Check if the user exists
	user, err := repository.GetUserByEmail(loginRequest.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("login failed: %w", ErrUserNotFound)
		} else {
			return "", fmt.Errorf("unexpected error: %w", err)
		}
	}

	// Check if the password is correct
	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(loginRequest.Password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			log.Infof("Incorrect password for user %v", user.Email)
			return "", fmt.Errorf("login failed: %w", ErrInvalidCredentials)
		} else {
			log.Errorf("Error comparing passwords: %v", err)
			return "", err
		}
	}

	log.Infof("User %v logged in", user.Email)

	// Generate a JWT token
	token, err := generateJWT(user)
	if err != nil {
		return "", fmt.Errorf("failed to generate JWT: %w", err)
	}

	return token, nil
}
