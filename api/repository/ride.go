package repository

import (
	"fmt"

	"github.com/RensVanGiersbergen/skatetracker/models"
)

func AddRide(ride models.Ride) (models.Ride, error) {
	var createdRide models.Ride

	// Execute the query and return the ride id
	err := db.QueryRow(queryStore["add_ride.sql"], ride.UserId, ride.BoardId, ride.Title, ride.Description).Scan(&createdRide.RideId, &createdRide.UserId, &createdRide.BoardId, &createdRide.Title, &createdRide.Description, &createdRide.StartTime)
	if err != nil {
		return models.Ride{}, fmt.Errorf("error executing query: %w", err)
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
