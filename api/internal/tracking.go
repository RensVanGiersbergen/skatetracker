package internal

import (
	"math"

	"github.com/RensVanGiersbergen/skatetracker/models"
)

// Haversine formula to calculate distance between two lat/lon points
func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371.0 // Earth radius in kilometers

	// Convert degrees to radians
	lat1Rad := lat1 * (math.Pi / 180)
	lat2Rad := lat2 * (math.Pi / 180)
	dLat := (lat2 - lat1) * (math.Pi / 180)
	dLon := (lon2 - lon1) * (math.Pi / 180)

	// Haversine calculation
	a := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return R * c * 1000 // Distance in meters
}

// Calculate total distance from a list of trackings
func CalculateTotalDistance(trackings []models.Tracking) uint32 {
	totalDistance := 0.0

	// Iterate through consecutive points
	for i := 1; i < len(trackings); i++ {
		lat1, lon1 := trackings[i-1].Latitude, trackings[i-1].Longitude
		lat2, lon2 := trackings[i].Latitude, trackings[i].Longitude
		totalDistance += haversine(lat1, lon1, lat2, lon2)
	}

	roundedDistance := uint32(math.Round(totalDistance))
	return roundedDistance // Distance in meters
}

// Get top speed from a list of trackings
func GetTopSpeed(trackings []models.Tracking) float32 {
	topSpeed := float32(0)
	for _, tracking := range trackings {
		if tracking.Speed > topSpeed {
			topSpeed = tracking.Speed
		}
	}
	return topSpeed // Speed in m/s
}
