package repository

import (
	"database/sql"
	"time"

	"github.com/mbarolo/test_back/models"
	"github.com/mbarolo/test_back/utils"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) ExistsByEmail(email string) (bool, error) {
	var count int
	query := "SELECT COUNT(*) FROM " + TableNameUser + " WHERE email = ?"

	err := r.db.QueryRow(query, email).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *UserRepository) GetAll() ([]*models.User, error) {
	query := "SELECT * FROM  " + TableNameUser
	users, err := utils.GenericScanAll[models.User](r.db, query)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) GetById(id int64) (*models.User, error) {
	query := "SELECT * FROM " + TableNameUser + " WHERE id = ?"
	user, err := utils.GenericScanAll[models.User](r.db, query, id)
	if err != nil {
		return nil, err
	}
	if len(user) == 0 {
		return nil, sql.ErrNoRows
	}

	return user[0], nil
}

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	query := "SELECT * FROM " + TableNameUser + " WHERE email = ?"
	user, err := utils.GenericScanAll[models.User](r.db, query, email)
	if err != nil {
		return nil, err
	}
	if len(user) == 0 {
		return nil, sql.ErrNoRows
	}

	return user[0], nil
}

func (r *UserRepository) Create(user *models.User) (int64, error) {
	query := "INSERT INTO " + TableNameUser + " (email, hashed_password, first_name, last_name, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)"
	res, err := r.db.Exec(query, user.Email, user.HashedPassword, user.FirstName, user.LastName, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

func (r *UserRepository) Update(user *models.User) (int64, error) {
	query := "UPDATE " + TableNameUser + " SET email = ?, hashed_password = ?, first_name = ?, last_name = ?, updated_at = ?"
	res, err := r.db.Exec(query, user.Email, user.HashedPassword, user.FirstName, user.LastName, time.Now())
	if err != nil {
		return -1, err
	}

	return res.RowsAffected()
}

func (r *UserRepository) Delete(id string) (int64, error) {
	query := "DELETE FROM " + TableNameUser + " WHERE id = ?"
	res, err := r.db.Exec(query)
	if err != nil {
		return -1, err
	}

	return res.RowsAffected()
}
