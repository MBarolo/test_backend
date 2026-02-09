package repository

import (
	"database/sql"
	"time"

	"github.com/mbarolo/test_back/models"
	"github.com/mbarolo/test_back/utils"
)

type BikeRepository struct {
	db *sql.DB
}

func NewBikeRepository(db *sql.DB) *BikeRepository {
	return &BikeRepository{db}
}

func (r *BikeRepository) IsAvailable(id string) (bool, error) {
	var available bool
	query := "SELECT available FROM " + TableNameBike + " WHERE id = ?"
	err := r.db.QueryRow(query, id).Scan(&available)
	if err != nil {
		return false, err
	}

	return available, nil
}

func (r *BikeRepository) GetAll() ([]*models.Bike, error) {
	query := "SELECT * FROM " + TableNameBike
	bikes, err := utils.GenericScanAll[models.Bike](r.db, query)
	if err != nil {
		return nil, err
	}

	return bikes, nil
}

func (r *BikeRepository) GetAllAvailable() ([]*models.Bike, error) {
	query := "SELECT * FROM " + TableNameBike + " WHERE is_available = 1"
	bikes, err := utils.GenericScanAll[models.Bike](r.db, query)
	if err != nil {
		return nil, err
	}

	return bikes, nil
}

func (r *BikeRepository) GetById(id int64) (*models.Bike, error) {
	query := "SELECT * FROM " + TableNameBike + " WHERE id = ?"
	bike, err := utils.GenericScanAll[models.Bike](r.db, query, id)
	if err != nil {
		return nil, err
	}
	if len(bike) == 0 {
		return nil, sql.ErrNoRows
	}
	return bike[0], nil
}

func (r *BikeRepository) CreateBike(bike *models.Bike) (int64, error) {
	query := "INSERT INTO " + TableNameBike + " (is_available, latitude, longitude, cost_per_minute, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)"
	res, err := r.db.Exec(query, bike.IsAvailable, bike.Latitude, bike.Longitude, bike.CostPerMinute, bike.CreatedAt, bike.UpdatedAt)
	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

func (r *BikeRepository) UpdateBike(bike *models.Bike) (int64, error) {
	query := "UPDATE " + TableNameBike + " SET is_available = ?, latitude = ?, longitude = ?, updated_at = ?"
	res, err := r.db.Exec(query, bike.IsAvailable, bike.Latitude, bike.Longitude, time.Now())
	if err != nil {
		return -1, nil
	}
	return res.RowsAffected()
}
