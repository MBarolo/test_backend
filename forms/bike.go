package forms

import "github.com/mbarolo/test_back/models"

type BikeForm struct {
	IsAvailable   *bool    `json:"is_available"`
	Latitude      *float64 `json:"latitude"`
	Longitude     *float64 `json:"longitude"`
	CostPerMinute *int     `json:"cost_per_minute"`
}

func (bf *BikeForm) ToBike() *models.Bike {
	return &models.Bike{
		IsAvailable:   *bf.IsAvailable,
		Latitude:      *bf.Latitude,
		Longitude:     *bf.Longitude,
		CostPerMinute: *bf.CostPerMinute,
	}
}
