package repository

import (
	"fmt"

	"github.com/RensVanGiersbergen/skatetracker/models"
)

func AddBoard(board models.AddBoard) (models.Board, error) {
	var newBoard models.Board

	// Add board
	err := db.QueryRow(queryStore["add_board.sql"], board.UserId, board.Nickname, board.Brand, board.PrimaryBoard).Scan(&newBoard.BoardId, &newBoard.UserId, &newBoard.Nickname, &newBoard.Brand, &newBoard.RideCount, &newBoard.TotalDistance, &newBoard.TopSpeed, &newBoard.TotalRidetime, &newBoard.PrimaryBoard, &newBoard.CreatedAt)
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

func GetBoardsByUser(userID string) ([]models.Board, error) {
	var boards []models.Board

	// Get boards by user (returns primary board first)
	rows, err := db.Query(queryStore["get_boards_by_user.sql"], userID)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var board models.Board
		err := rows.Scan(&board.BoardId, &board.UserId, &board.Nickname, &board.Brand, &board.RideCount, &board.TotalDistance, &board.TopSpeed, &board.TotalRidetime, &board.PrimaryBoard, &board.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		boards = append(boards, board)
	}

	return boards, nil
}

func GetBoardById(boardId string) (models.Board, error) {
	var board models.Board

	// Get board by id
	err := db.QueryRow(queryStore["get_board_by_id.sql"], boardId).Scan(&board.BoardId, &board.UserId, &board.Nickname, &board.Brand, &board.RideCount, &board.TotalDistance, &board.TopSpeed, &board.TotalRidetime, &board.PrimaryBoard, &board.CreatedAt)
	if err != nil {
		// Check if error is no rows found
		if err.Error() == "sql: no rows in result set" {
			return models.Board{}, fmt.Errorf("board not found")
		} else {
			return models.Board{}, fmt.Errorf("error executing query: %w", err)
		}
	}

	return board, nil
}

func UpdateBoard(board models.Board) (models.Board, error) {
	var updatedBoard models.Board

	// Update board
	err := db.QueryRow(queryStore["update_board.sql"], board.BoardId, board.Nickname, board.Brand, board.PrimaryBoard).Scan(&updatedBoard.BoardId, &updatedBoard.UserId, &updatedBoard.Nickname, &updatedBoard.Brand, &updatedBoard.RideCount, &updatedBoard.TotalDistance, &updatedBoard.TopSpeed, &updatedBoard.TotalRidetime, &updatedBoard.PrimaryBoard, &updatedBoard.CreatedAt)
	if err != nil {
		return models.Board{}, fmt.Errorf("error executing query: %w", err)
	}

	return updatedBoard, nil
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
