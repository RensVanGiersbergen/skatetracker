package models

// Model struct for ride
type AddRide struct {
	UserId      string     `json:"user_id" db:"user_id"`
	BoardId     string     `json:"board_id" db:"board_id" binding:"required"`
	Title       string     `json:"title" db:"title" binding:"required"`
	Description string     `json:"description,omitempty" db:"description"`
	Trackings   []Tracking `json:"trackings" binding:"required"`
}
