package repository

import (
	"fmt"

	"github.com/RensVanGiersbergen/skatetracker/models"
)

func AddTrackingsToRide(trackings []models.Tracking) error {
	// Add the trackings to the database
	for _, tracking := range trackings {
		_, err := db.Exec(queryStore["add_tracking.sql"], tracking.RideId, tracking.TrackingTime, tracking.Latitude, tracking.Longitude, tracking.Speed, tracking.Shakiness)
		if err != nil {
			return fmt.Errorf("error executing query: %w", err)
		}
	}
	return nil
}

func GetTrackingsForRide(rideId string) ([]models.Tracking, error) {
	rows, err := db.Query(queryStore["get_trackings_by_ride.sql"], rideId)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()

	trackings := make([]models.Tracking, 0)
	for rows.Next() {
		var tracking models.Tracking
		err = rows.Scan(&tracking.RideId, &tracking.TrackingTime, &tracking.Latitude, &tracking.Longitude, &tracking.Speed, &tracking.Shakiness)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		trackings = append(trackings, tracking)
	}
	return trackings, nil
}
