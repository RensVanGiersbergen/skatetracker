package models

import "time"

// Model struct for board
type Board struct {
	BoardId       string    `json:"board_id" db:"board_id"`
	UserId        string    `json:"user_id" db:"user_id"`
	Nickname      string    `json:"nickname" db:"nickname"`
	Brand         string    `json:"brand" db:"brand"`
	RideCount     int       `json:"ride_count" db:"ride_count"`
	TotalDistance uint      `json:"total_distance" db:"total_distance"`
	TopSpeed      float32   `json:"top_speed" db:"top_speed"`
	TotalRidetime uint      `json:"total_ridetime" db:"total_ridetime"` // in seconds
	PrimaryBoard  bool      `json:"primary_board" db:"primary_board"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
}
