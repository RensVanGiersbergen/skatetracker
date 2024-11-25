package internal

import (
	"errors"
	"fmt"
	"strings"

	"github.com/RensVanGiersbergen/skatetracker/models"
	"github.com/RensVanGiersbergen/skatetracker/repository"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

var (
	// Errors
	ErrInvalidEmail      = errors.New("invalid email")
	ErrDuplicateEmail    = errors.New("email already in use")
	ErrDuplicateUsername = errors.New("username already in use")
)

func Register(register models.RegisterRequest) error {
	// Check if email has @
	if !strings.Contains(register.Email, "@") {
		log.Debugf("invalid email: %v", register.Email)
		return fmt.Errorf("invalid email: %w", ErrInvalidEmail)
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Errorf("failed to hash password: %v", err)
		return fmt.Errorf("failed to hash password: %w", err)
	}

	// Add the user to the database
	err = repository.AddUser(models.RegisterRequest{
		Username: register.Username,
		Email:    register.Email,
		Password: string(hashedPassword),
	})
	if err != nil {
		switch {
		case err.Error() == "username already exists":
			log.Debugf("username already in use: %v", register.Username)
			return fmt.Errorf("username already in use: %w", ErrDuplicateUsername)
		case err.Error() == "email already exists":
			log.Debugf("email already in use: %v", register.Email)
			return fmt.Errorf("email already in use: %w", ErrDuplicateEmail)
		default:
			log.Errorf("error adding user: %v", err)
			return fmt.Errorf("error adding user: %w", err)
		}
	}

	return nil
}
