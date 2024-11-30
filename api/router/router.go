package router

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/RensVanGiersbergen/skatetracker/internal"
	"github.com/RensVanGiersbergen/skatetracker/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Define route prefixes
	accountPrefix := "/account"
	ridePrefix := "/ride"

	//Define all endpoints
	// Unprotected endpoints
	// Login
	router.POST(accountPrefix+"/login", func(c *gin.Context) {
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

	// Register
	router.POST(accountPrefix+"/register", func(c *gin.Context) {
		var registerRequest models.RegisterRequest

		// Bind the request body to the registerRequest struct
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

	// Use middleware for protected endpoints from here on
	router.Use(authMiddleware())

	// Protected endpoints (JWT required)
	// Account endpoints
	// Verify (JWT)
	router.GET("/account/verify", func(c *gin.Context) {
		claims := c.MustGet("claims").(jwt.MapClaims)
		c.JSON(http.StatusOK, gin.H{"message": "user is verified", "claims": claims})
	})

	// Ride endpoints
	// Add ride
	router.POST(ridePrefix+`/add`, func(c *gin.Context) {
		var ride models.Ride

		// Bind the request body to the ride struct
		if err := c.ShouldBindJSON(&ride); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
			return
		}

		// Get the user id from the claims and add it to the ride struct
		claims := c.MustGet("claims").(jwt.MapClaims)
		ride.UserId = claims["sub"].(string)

		// Add the ride
		createdRide, err := internal.AddRide(ride)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}

		c.JSON(http.StatusOK, createdRide)
	})

	// Finish ride
	router.PUT(ridePrefix+`/finish`, func(c *gin.Context) {
		var ride models.Ride

		// Bind the request body to the ride struct
		if err := c.ShouldBindJSON(&ride); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
			return
		}

		// Get the user id from the claims and add it to the ride struct
		claims := c.MustGet("claims").(jwt.MapClaims)
		ride.UserId = claims["sub"].(string)

		// Finish the ride
		finishedRide, err := internal.FinishRide(ride)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrNotOwnerOfRide):
				c.JSON(http.StatusUnauthorized, gin.H{"error": "user is not the owner of the ride"})
			case errors.Is(err, internal.ErrNotEnoughTrackings):
				c.JSON(http.StatusBadRequest, gin.H{"error": "not enough trackings to finish ride"})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			}
			return
		}

		c.JSON(http.StatusOK, finishedRide)
	})

	// Delete ride
	router.DELETE(ridePrefix+`/delete/:rideId`, func(c *gin.Context) {
		rideId := c.Param("rideId")

		// Get the user id from the claims
		claims := c.MustGet("claims").(jwt.MapClaims)
		userId := claims["sub"].(string)

		// Delete the ride
		err := internal.DeleteRide(userId, rideId)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrRideNotFound):
				c.JSON(http.StatusNotFound, gin.H{"error": "ride not found"})
			case errors.Is(err, internal.ErrNotOwnerOfRide):
				c.JSON(http.StatusUnauthorized, gin.H{"error": "user is not the owner of the ride"})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			}
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "ride deleted"})
	})

	// Get all rides by user with pagination
	router.GET(ridePrefix+`/all`, func(c *gin.Context) {
		// Get the user id from the claims
		claims := c.MustGet("claims").(jwt.MapClaims)
		userId := claims["sub"].(string)

		// Get pagination parameters from query
		page := c.DefaultQuery("page", "1")
		limit := c.DefaultQuery("limit", "10")

		// Convert limit and page to integers
		limitInt, err := strconv.Atoi(limit)
		if err != nil || limitInt <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit"})
			return
		}

		pageInt, err := strconv.Atoi(page)
		if err != nil || pageInt <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid page"})
			return
		}

		// Get the rides
		rides, err := internal.GetAllRidesByUserWithPagination(userId, pageInt, limitInt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}

		c.JSON(http.StatusOK, rides)
	})

	// Get ride by id
	router.GET(ridePrefix+`/:rideId`, func(c *gin.Context) {
		rideId := c.Param("rideId")

		var ride models.Ride

		// Get the user id from the claims
		claims := c.MustGet("claims").(jwt.MapClaims)
		userId := claims["sub"].(string)

		// Get the ride
		ride, err := internal.GetRideById(userId, rideId)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrRideNotFound):
				c.JSON(http.StatusNotFound, gin.H{"error": "ride not found"})
			case errors.Is(err, internal.ErrNotOwnerOfRide):
				c.JSON(http.StatusUnauthorized, gin.H{"error": "user is not the owner of the ride"})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			}
			return
		}

		c.JSON(http.StatusOK, ride)
	})
	return router
}
