package forms

import (
	"time"

	"github.com/mbarolo/test_back/models"
)

type StartEndRentalForm struct {
	BikeID int64 `json:"bike_id"`
}

type RentalForm struct {
	UserID         *int64              `json:"user_id"`
	BikeID         *int64              `json:"bike_id"`
	Status         models.RentalStatus `json:"rental_status"`
	StartTime      *time.Time          `json:"start_time"`
	EndTime        *time.Time          `json:"end_time"`
	StartLatitude  *float64            `json:"start_latitude"`
	StartLongitude *float64            `json:"start_longitude"`
	EndLatitude    *float64            `json:"end_latitude"`
	EndLongitude   *float64            `json:"end_longitude"`
	Duration       *int                `json:"duration"` // minutes
}

/*
func (rf *RentalForm) ToRental() *models.Rental {
	return &models.Rental{
		UserID:         rf.UserID,
		BikeID:         rf.BikeID,
		Status:         rf.Status,
		StartTime:      rf.StartTime,
		EndTime:        rf.EndTime,
		StartLatitude:  rf.StartLatitude,
		StartLongitude: *rf.StartLongitude,
		EndLatitude:    *rf.EndLatitude,
		EndLongitude:   *rf.EndLatitude,
		Duration:       *rf.Duration,
	}
}
*/
