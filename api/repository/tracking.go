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
