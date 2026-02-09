package models

import (
	"errors"
	"time"
)

type Bike struct {
	Id            int64     `json:"id"`
	IsAvailable   bool      `json:"is_available"`
	Latitude      float64   `json:"latitude"`
	Longitude     float64   `json:"longitude"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	CostPerMinute int       `json:"cost_per_minute"`
}

func (b *Bike) ValidateFields() error {
	if b.CostPerMinute < 0 {
		return errors.New("costo inválido")
	}

	return nil
}
