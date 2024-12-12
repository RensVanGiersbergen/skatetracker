package internal

import (
	"fmt"

	"github.com/RensVanGiersbergen/skatetracker/models"
	"github.com/RensVanGiersbergen/skatetracker/repository"
	log "github.com/sirupsen/logrus"
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
