package internal

import (
	"errors"
	"fmt"

	"github.com/RensVanGiersbergen/skatetracker/models"
	"github.com/RensVanGiersbergen/skatetracker/repository"
	log "github.com/sirupsen/logrus"
)

var (
	// Errors
	ErrBoardNotFound   = errors.New("board not found")
	ErrNotOwnerOfBoard = errors.New("user is not the owner of the board")
)

func AddBoard(board models.AddBoard) (models.Board, error) {
	// Add board
	addedBoard, err := repository.AddBoard(board)
	if err != nil {
		log.Errorf("error adding board: %v", err)
		return models.Board{}, fmt.Errorf("error adding board: %w", err)
	}
	return addedBoard, nil
}

func GetBoardsByUser(userID string) ([]models.Board, error) {
	// Get boards by user
	boards, err := repository.GetBoardsByUser(userID)
	if err != nil {
		log.Errorf("error getting boards by user: %v", err)
		return nil, fmt.Errorf("error getting boards by user: %w", err)
	}
	return boards, nil
}

func GetBoardById(boardId string) (models.Board, error) {
	// Get board by id
	board, err := repository.GetBoardById(boardId)
	if err != nil {
		if err.Error() == "board not found" {
			return models.Board{}, ErrBoardNotFound
		} else {
			log.Errorf("error getting board by id: %v", err)
			return models.Board{}, fmt.Errorf("error getting board by id: %w", err)
		}
	}
	return board, nil
}

func UpdateBoard(board models.Board) (models.Board, error) {
	// Get the board by id
	currentBoard, err := repository.GetBoardById(board.BoardId)
	if err != nil {
		if err.Error() == "board not found" {
			return models.Board{}, ErrBoardNotFound
		} else {
			log.Errorf("error getting board by id: %v", err)
			return models.Board{}, fmt.Errorf("error getting board by id: %w", err)
		}
	}

	// Check if the user is the owner of the board
	if currentBoard.UserId != board.UserId {
		return models.Board{}, ErrNotOwnerOfBoard
	}

	updatedBoard, err := repository.UpdateBoard(board)
	if err != nil {
		log.Errorf("error updating board: %v", err)
		return models.Board{}, fmt.Errorf("error updating board: %w", err)
	}

	return updatedBoard, nil
}
