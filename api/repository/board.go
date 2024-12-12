package repository

import (
	"fmt"

	"github.com/RensVanGiersbergen/skatetracker/models"
)

func AddBoard(board models.AddBoard) (models.Board, error) {
	var newBoard models.Board

	// Add board
	err := db.QueryRow(queryStore["add_board.sql"], board.UserId, board.Nickname, board.Brand).Scan(&newBoard.BoardId, &newBoard.UserId, &newBoard.Nickname, &newBoard.Brand, &newBoard.RideCount, &newBoard.TotalDistance, &newBoard.TopSpeed, &newBoard.TotalRidetime, &newBoard.PrimaryBoard, &newBoard.CreatedAt)
	if err != nil {
		return models.Board{}, fmt.Errorf("error executing query: %w", err)
	}

	// Add board default stats
	newBoard.RideCount = 0
	newBoard.TotalDistance = 0
	newBoard.TopSpeed = 0.0
	newBoard.TotalRidetime = 0

	return newBoard, nil
}

// updateBoardStats updates the board statistics based on a ride added
func updateBoardStats(ride models.Ride) error {
	// Calculate total seconds of the ride
	totalSeconds := ride.EndTime.Sub(ride.StartTime).Seconds()

	_, err := db.Exec(queryStore["update_board_stats.sql"], ride.BoardId, ride.Distance, ride.TopSpeed, totalSeconds)
	if err != nil {
		return fmt.Errorf("error executing query: %w", err)
	}
	return nil
}
