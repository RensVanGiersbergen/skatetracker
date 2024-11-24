package internal

import (
	"os"
	"time"

	"github.com/RensVanGiersbergen/skatetracker/models"
	"github.com/dgrijalva/jwt-go"
)

func generateJWT(user models.User) (string, error) {
	// Create a new token object, specifying signing method and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":      user.UserId,
		"username": user.Username,
		"email":    user.Email,
		"exp":      time.Now().AddDate(1, 0, 0).Unix(),
	})

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
