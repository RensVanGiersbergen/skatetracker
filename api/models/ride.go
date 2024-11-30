package models

import "time"

// Model struct for ride
type Ride struct {
	RideId      string     `json:"ride_id" db:"ride_id"`
	UserId      string     `json:"user_id" db:"user_id"`
	BoardId     string     `json:"board_id" db:"board_id" binding:"required"`
	Completed   bool       `json:"completed" db:"completed"`
	Title       string     `json:"title" db:"title" binding:"required"`
	Description string     `json:"description,omitempty" db:"description"`
	StartTime   time.Time  `json:"start_time" db:"start_time"`
	EndTime     *time.Time `json:"end_time,omitempty" db:"end_time"`
	Distance    *uint32    `json:"distance,omitempty" db:"distance"`
	TopSpeed    *float32   `json:"top_speed,omitempty" db:"top_speed"`
	Trackings   []Tracking `json:"trackings,omitempty"`
}
