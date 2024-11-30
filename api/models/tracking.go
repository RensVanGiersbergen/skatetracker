package models

import "time"

type Tracking struct {
	TrackingId   int       `json:"tracking_id,omitempty" db:"tracking_id"`
	RideId       string    `json:"ride_id" db:"ride_id"`
	TrackingTime time.Time `json:"tracking_time" db:"tracking_time" binding:"required"`
	Latitude     float64   `json:"latitude" db:"latitude" binding:"required"`
	Longitude    float64   `json:"longitude" db:"longitude" binding:"required"`
	Speed        float32   `json:"speed" db:"speed" binding:"required"`
	Shakiness    float32   `json:"shakiness" db:"shakiness" binding:"required"`
}
