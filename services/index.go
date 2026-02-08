package services

import (
	"github.com/mbarolo/test_back/config"
	"github.com/mbarolo/test_back/repository"
)

var sqliteConnection = config.NewSQLiteConnection()

var (
	userRepo = repository.NewUserRepository(sqliteConnection.DB)
	// bikeRepo   = repository.NewBikeRepository(sqliteConnection.DB)
	// rentalRepo = repository.NewRentalRepository(sqliteConnection.DB)
)
