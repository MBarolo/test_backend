package models

import "time"

type RentalStatus string

const (
	RUNNING RentalStatus = "running"
	ENDED   RentalStatus = "ended"
)

type Rental struct {
	Id             int64        `json:"id"`
	UserId         int64        `json:"user_id"`
	BikeId         int64        `json:"bike_id"`
	RentalStatus   RentalStatus `json:"rental_status"`
	StartTime      time.Time    `json:"start_time"`
	EndTime        *time.Time   `json:"end_time"`
	StartLatitude  float64      `json:"start_latitude"`
	StartLongitude float64      `json:"start_longitude"`
	EndLatitude    *float64     `json:"end_latitude"`
	EndLongitude   *float64     `json:"end_longitude"`
	Duration       *int         `json:"duration"` // minutes
	Cost           *int         `json:"cost"`
}
