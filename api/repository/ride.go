package repository

import (
	"fmt"

	"github.com/RensVanGiersbergen/skatetracker/models"
)

func AddRide(ride models.Ride) (models.Ride, error) {
	var createdRide models.Ride

	// Execute the query and return the ride id
	err := db.QueryRow(queryStore["add_ride.sql"], ride.UserId, ride.BoardId, ride.Completed, ride.Title, ride.Description, ride.StartTime, ride.EndTime, ride.Distance, ride.TopSpeed).Scan(&createdRide.RideId, &createdRide.UserId, &createdRide.BoardId, &createdRide.Completed, &createdRide.Title, &createdRide.Description, &createdRide.StartTime, &createdRide.EndTime, &createdRide.Distance, &createdRide.TopSpeed)
	if err != nil {
		// Check if error is fk constraint error
		if err.Error() == "pq: insert or update on table \"rides\" violates foreign key constraint \"rides_board_id_fkey\"" {
			return models.Ride{}, fmt.Errorf("board not found")
		} else {
			return models.Ride{}, fmt.Errorf("error executing query: %w", err)
		}
	}

	// Set ride id for trackings
	for i := range ride.Trackings {
		ride.Trackings[i].RideId = createdRide.RideId
	}

	err1 := AddTrackingsToRide(ride.Trackings)
	if err1 != nil {
		return models.Ride{}, fmt.Errorf("error adding trackings to ride: %w", err1)
	}

	err2 := updateBoardStats(createdRide)
	if err2 != nil {
		return models.Ride{}, fmt.Errorf("error updating board stats: %w", err2)
	}
	return createdRide, nil
}

func GetAllRidesByUserWithPagination(userId string, page int, limit int) ([]models.Ride, error) {
	var rides []models.Ride

	rows, err := db.Query(queryStore["get_all_rides_by_user_with_pagination.sql"], userId, limit, (page-1)*limit)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var ride models.Ride
		err := rows.Scan(&ride.RideId, &ride.UserId, &ride.BoardId, &ride.Completed, &ride.Title, &ride.Description, &ride.StartTime, &ride.EndTime, &ride.Distance, &ride.TopSpeed)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		rides = append(rides, ride)
	}
	return rides, nil
}

func UpdateRide(ride models.Ride) (models.Ride, error) {
	var updatedRide models.Ride
	err := db.QueryRow(queryStore["update_ride.sql"], ride.RideId, ride.Completed, ride.Title, ride.Description, ride.StartTime, ride.EndTime, ride.Distance, ride.TopSpeed).Scan(&updatedRide.RideId, &updatedRide.UserId, &updatedRide.BoardId, &updatedRide.Completed, &updatedRide.Title, &updatedRide.Description, &updatedRide.StartTime, &updatedRide.EndTime, &updatedRide.Distance, &updatedRide.TopSpeed)
	if err != nil {
		return models.Ride{}, fmt.Errorf("error executing query: %w", err)
	}
	return updatedRide, nil
}

func GetRideById(rideId string) (models.Ride, error) {
	var ride models.Ride
	err := db.QueryRow(queryStore["get_ride_by_id.sql"], rideId).Scan(&ride.RideId, &ride.UserId, &ride.BoardId, &ride.Completed, &ride.Title, &ride.Description, &ride.StartTime, &ride.EndTime, &ride.Distance, &ride.TopSpeed)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return models.Ride{}, fmt.Errorf("ride not found")
		}
		return models.Ride{}, fmt.Errorf("error executing query: %w", err)
	}
	return ride, nil
}

func DeleteRide(rideId string) error {
	_, err := db.Exec(queryStore["delete_ride.sql"], rideId)
	if err != nil {
		return fmt.Errorf("error executing query: %w", err)
	}
	return nil
}
