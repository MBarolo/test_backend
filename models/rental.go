package models

import "time"

type RentalStatus string

const (
	RUNNING RentalStatus = "running"
	ENDED   RentalStatus = "ended"
)

type Rental struct {
	ID             int64        `json:"id"`
	UserID         int64        `json:"user_id"`
	BikeID         int64        `json:"bike_id"`
	Status         RentalStatus `json:"rental_status"`
	StartTime      time.Time    `json:"start_time"`
	EndTime        time.Time    `json:"end_time"`
	StartLatitude  float64      `json:"start_latitude"`
	StartLongitude float64      `json:"start_longitude"`
	EndLatitude    float64      `json:"end_latitude"`
	EndLongitude   float64      `json:"end_longitude"`
}
