package internal

import (
	"database/sql"
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
	ErrInvalidEmail       = errors.New("invalid email")
	ErrDuplicateEmail     = errors.New("email already in use")
	ErrDuplicateUsername  = errors.New("username already in use")
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
			log.Debugf("username already in use when registering: %v", register.Username)
			return fmt.Errorf("username already in use: %w", ErrDuplicateUsername)
		case err.Error() == "email already exists":
			log.Debugf("email already in use when registering: %v", register.Email)
			return fmt.Errorf("email already in use: %w", ErrDuplicateEmail)
		default:
			log.Errorf("error adding user: %v", err)
			return fmt.Errorf("error adding user: %w", err)
		}
	}

	return nil
}
