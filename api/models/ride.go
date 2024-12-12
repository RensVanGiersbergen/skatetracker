package models

import "time"

// Model struct for ride
type Ride struct {
	RideId      string     `json:"ride_id" db:"ride_id"`
	UserId      string     `json:"user_id" db:"user_id"`
	BoardId     string     `json:"board_id" db:"board_id"`
	Completed   bool       `json:"completed" db:"completed"`
	Title       string     `json:"title" db:"title"`
	Description string     `json:"description,omitempty" db:"description"`
	StartTime   time.Time  `json:"start_time" db:"start_time"`
	EndTime     time.Time  `json:"end_time" db:"end_time"`
	Distance    uint32     `json:"distance" db:"distance"`
	TopSpeed    float32    `json:"top_speed" db:"top_speed"`
	Trackings   []Tracking `json:"trackings,omitempty"`
}
