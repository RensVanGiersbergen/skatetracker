package router

import (
	"errors"
	"net/http"

	"github.com/RensVanGiersbergen/skatetracker/internal"
	"github.com/RensVanGiersbergen/skatetracker/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	//Define all endpoints
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, "chill")
	})

	// Unprotected endpoints

	router.POST("/account/login", func(c *gin.Context) {
		var loginRequest models.LoginRequest

		// Bind the request body to the LoginRequest struct
		if err := c.ShouldBindJSON(&loginRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
			return
		}

		// Check if the user exists
		token, err := internal.Login(loginRequest)
		if err != nil {
			// Check if the error is a custom error
			switch {
			case errors.Is(err, internal.ErrUserNotFound):
				c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			case errors.Is(err, internal.ErrInvalidCredentials):
				c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			}
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token, "message": "user logged in"})
	})

	router.POST("/account/register", func(c *gin.Context) {
		var registerRequest models.RegisterRequest

		// Bind the request body to the LoginRequest struct
		if err := c.ShouldBindJSON(&registerRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
			return
		}

		// Register the user
		err := internal.Register(registerRequest)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrInvalidEmail):
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid email"})
			case errors.Is(err, internal.ErrDuplicateUsername):
				c.JSON(http.StatusBadRequest, gin.H{"error": "username already in use"})
			case errors.Is(err, internal.ErrDuplicateEmail):
				c.JSON(http.StatusBadRequest, gin.H{"error": "email already in use"})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			}
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "user registered"})

	})

	router.Use(authMiddleware())

	// Protected endpoints (JWT required)
	// Account endpoints
	router.GET("/account/verify", func(c *gin.Context) {
		claims := c.MustGet("claims").(jwt.MapClaims)
		c.JSON(http.StatusOK, gin.H{"message": "user is verified", "claims": claims})
	})

	return router
}
