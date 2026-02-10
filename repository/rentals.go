package repository

import (
	"database/sql"

	"github.com/mbarolo/test_back/models"
	"github.com/mbarolo/test_back/utils"
)

type RentalRepository struct {
	db *sql.DB
}

func NewRentalRepository(db *sql.DB) *RentalRepository {
	return &RentalRepository{db}
}

func (r *RentalRepository) GetAll() ([]*models.Rental, error) {
	query := "SELECT * FROM " + TableNameRental
	rentals, err := utils.GenericScanAll[models.Rental](r.db, query)
	if err != nil {
		return nil, err
	}
	if len(rentals) == 0 {
		return []*models.Rental{}, nil
	}
	return rentals, nil
}

func (r *RentalRepository) GetUserHistory(userId int64) ([]*models.Rental, error) {
	query := "SELECT * FROM " + TableNameRental + " WHERE user_id = ?"
	rentals, err := utils.GenericScanAll[models.Rental](r.db, query, userId)
	if err != nil {
		return nil, err
	}

	return rentals, nil
}

func (r *RentalRepository) GetById(id int64) (*models.Rental, error) {
	query := "SELECT * FROM " + TableNameRental + " WHERE id = ?"
	rental, err := utils.GenericScanAll[models.Rental](r.db, query, id)
	if err != nil {
		return nil, err
	}
	if len(rental) == 0 {
		return nil, sql.ErrNoRows
	}
	return rental[0], nil
}

func (r *RentalRepository) GetRunningRental(userId int64) (*models.Rental, error) {
	query := "SELECT * FROM " + TableNameRental + " WHERE user_id = ? AND rental_status = ?"
	rental, err := utils.GenericScanAll[models.Rental](r.db, query, userId, models.RUNNING)
	if err != nil {
		return nil, err
	}
	if len(rental) == 0 {
		return nil, nil
	}
	return rental[0], nil
}

func (r *RentalRepository) Create(rental *models.Rental) (int64, error) {
	query := "INSERT INTO " + TableNameRental + " (user_id, bike_id, rental_status, start_time, end_time, start_latitude, start_longitude, end_latitude, end_longitude) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"
	res, err := r.db.Exec(query, rental.UserId, rental.BikeId, rental.RentalStatus, rental.StartTime, rental.EndTime, rental.StartLatitude, rental.StartLongitude, rental.EndLatitude, rental.EndLongitude)
	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

func (r *RentalRepository) Update(rental *models.Rental) (int64, error) {
	query := "UPDATE " + TableNameRental + " SET user_id = ?, bike_id = ?, rental_status = ?, start_time = ?, end_time = ?, start_latitude = ?, start_longitude = ?, end_latitude = ?, end_longitude = ?, duration = ?, cost = ? WHERE id = ?"
	res, err := r.db.Exec(query, rental.UserId, rental.BikeId, rental.RentalStatus, rental.StartTime, rental.EndTime, rental.StartLatitude, rental.StartLongitude, rental.EndLatitude, rental.EndLongitude, rental.Duration, rental.Cost, rental.Id)
	if err != nil {
		return -1, err
	}

	return res.RowsAffected()
}
