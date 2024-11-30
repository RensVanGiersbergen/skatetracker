package internal

import (
	"errors"
	"fmt"

	"github.com/RensVanGiersbergen/skatetracker/models"
	"github.com/RensVanGiersbergen/skatetracker/repository"
	log "github.com/sirupsen/logrus"
)

var (
	ErrNotEnoughTrackings = errors.New("not enough trackings to finish ride")
	ErrNotOwnerOfRide     = errors.New("user is not the owner of the ride")
	ErrRideNotFound       = errors.New("ride not found")
)

func GetAllRidesByUserWithPagination(userId string, page int, limit int) ([]models.Ride, error) {
	// Get all rides by the user with pagination
	rides, err := repository.GetAllRidesByUserWithPagination(userId, page, limit)
	if err != nil {
		log.Errorf("Failed to get all rides by user with pagination: %v", err)
		return nil, err
	}
	return rides, nil
}

func GetRideById(userId, rideId string) (models.Ride, error) {
	// Get the ride by the id
	ride, err := repository.GetRideById(rideId)
	if err != nil {
		if err.Error() == "ride not found" {
			return models.Ride{}, ErrRideNotFound
		} else {
			log.Errorf("Failed to get ride by id: %v", err)
			return models.Ride{}, fmt.Errorf("failed to get ride by id: %w", err)
		}
	}

	// Check if the owner of the ride is the user
	err1 := checkIfUserIsOwnerOfRide(ride.UserId, rideId)
	if err1 != nil {
		return models.Ride{}, err1
	}

	return ride, nil
}

func AddRide(ride models.Ride) (models.Ride, error) {
	// Add the ride to the database
	createdRide, err := repository.AddRide(ride)
	if err != nil {
		log.Errorf("Failed to add ride to the database: %v", err)
		return models.Ride{}, err
	}
	return createdRide, nil
}

func FinishRide(ride models.Ride) (models.Ride, error) {
	// Check if there are enough trackings to finish the ride
	if len(ride.Trackings) < 2 {
		return models.Ride{}, ErrNotEnoughTrackings
	}

	// Check if user is owner of the ride
	err := checkIfUserIsOwnerOfRide(ride.UserId, ride.RideId)
	if err != nil {
		if errors.Is(err, ErrNotOwnerOfRide) {
			return models.Ride{}, ErrNotOwnerOfRide
		} else {
			return models.Ride{}, err
		}
	}
	// Insert the trackings into the database
	err1 := repository.AddTrackingsToRide(ride.Trackings)
	if err1 != nil {
		log.Errorf("Failed to add trackings to ride: %v", err1)
		return models.Ride{}, fmt.Errorf("failed to add trackings to ride: %w", err1)
	}

	// Calculate the distance and top speed
	ride.Distance = CalculateTotalDistance(ride.Trackings)
	ride.TopSpeed = CalculateTopSpeed(ride.Trackings)

	// Set ride to finished
	ride.Completed = true
	ride.EndTime = &ride.Trackings[len(ride.Trackings)-1].TrackingTime

	// Update the ride to finish it
	updatedRide, err2 := repository.UpdateRide(ride)
	if err2 != nil {
		log.Errorf("Failed to update ride: %v", err2)
		return models.Ride{}, err2
	}

	// Set trackings to updated ride
	updatedRide.Trackings = ride.Trackings

	return updatedRide, nil
}

func DeleteRide(userId, rideId string) error {
	// Check if user is owner of the ride
	err := checkIfUserIsOwnerOfRide(userId, rideId)
	if err != nil {
		return err
	}
	// Delete the ride from the database
	err1 := repository.DeleteRide(rideId)
	if err1 != nil {
		log.Errorf("Failed to delete ride: %v", err1)
		return fmt.Errorf("failed to delete ride: %w", err1)
	}
	return nil
}

func checkIfUserIsOwnerOfRide(userId string, rideId string) error {
	// Get the ride from the database
	ride, err := repository.GetRideById(rideId)
	if err != nil {
		if err.Error() == "ride not found" {
			return ErrRideNotFound
		} else {
			log.Errorf("Failed to get ride by id: %v", err)
			return fmt.Errorf("failed to get ride by id: %w", err)
		}
	}

	// Check if the user is the owner of the ride
	if ride.UserId != userId {
		return ErrNotOwnerOfRide
	}
	return nil
}
