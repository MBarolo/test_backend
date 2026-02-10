package services

import (
	"database/sql"
	"errors"
	"time"

	"github.com/mbarolo/test_back/forms"
	"github.com/mbarolo/test_back/models"
	"github.com/mbarolo/test_back/utils"
)

func StartRental(currentUser *models.User, rental *forms.StartEndRentalForm) (*models.Rental, error) {
	bike, err := bikeRepo.GetById(rental.BikeID)
	if err != nil {
		return nil, errors.New("error al obtener la bicicleta: " + err.Error())
	}

	if !bike.IsAvailable {
		return nil, errors.New("bicicleta no disponible")
	}

	running, err := rentalRepo.GetRunningRental(currentUser.Id)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if running != nil {
		return nil, errors.New("usuario ya tiene un alquiler en curso")
	}

	newRental := models.Rental{
		UserId:         currentUser.Id,
		BikeId:         bike.Id,
		RentalStatus:   models.RUNNING,
		StartTime:      time.Now(),
		EndTime:        nil,
		StartLatitude:  bike.Latitude,
		StartLongitude: bike.Longitude,
	}

	newId, err := rentalRepo.Create(&newRental)
	if err != nil {
		return nil, err
	}
	newRental.Id = newId

	// Actualizamos la bicicleta a no disponible
	bike.IsAvailable = false
	_, err = bikeRepo.UpdateBike(bike)
	if err != nil {
		return nil, errors.New("error al actualizar la bicicleta: " + err.Error())
	}

	return &newRental, nil
}

func EndRental(currentUser *models.User, rental *forms.StartEndRentalForm) (*models.Rental, error) {
	bike, err := bikeRepo.GetById(rental.BikeID)
	if err != nil {
		return nil, err
	}

	running, err := rentalRepo.GetRunningRental(currentUser.Id)
	if err != nil {
		return nil, err
	}

	if running == nil {
		return nil, errors.New("el usuario no tiene un alquiler activo")
	}

	if running.BikeId != bike.Id {
		return nil, errors.New("el usuario no está alquilando esta bicicleta")
	}

	// Calculamos duracion del rental
	endTime := time.Now()
	running.EndTime = &endTime
	duration := int(endTime.Sub(running.StartTime).Minutes())
	running.Duration = &duration
	cost := bike.CostPerMinute * duration
	running.Cost = &cost

	// Asignamos end lat y long ~5km
	endlatitude, endLongitude := utils.GenerateRandomCoordinatesWithinRadius(running.StartLatitude, running.StartLongitude, 5.0)
	running.EndLatitude, running.EndLongitude = &endlatitude, &endLongitude

	running.RentalStatus = models.ENDED

	// Se actualiza bike y rental
	_, err = rentalRepo.Update(running)
	if err != nil {
		return nil, errors.New("error al actualizar el alquiler: " + err.Error())
	}

	bike.IsAvailable = true
	bike.Latitude = *running.EndLatitude
	bike.Longitude = *running.EndLongitude
	_, err = bikeRepo.UpdateBike(bike)
	if err != nil {
		return nil, errors.New("error al actualizar la bicicleta: " + err.Error())
	}

	return running, nil
}

func GetRentalHistory(userId int64) ([]*models.Rental, error) {
	rentals, err := rentalRepo.GetUserHistory(userId)
	if err != nil {
		return nil, err
	}

	return rentals, nil
}

func GetAllRentals() ([]*models.Rental, error) {
	rentals, err := rentalRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return rentals, err
}

func GetRentalById(id int64) (*models.Rental, error) {
	rental, err := rentalRepo.GetById(id)
	if err != nil {
		return nil, err
	}

	return rental, nil
}

func UpdateRental(id int64, updatedRental *forms.RentalForm) (*models.Rental, error) {
	originalRental, err := GetRentalById(id)
	if err != nil {
		return nil, err
	}

	if updatedRental.UserID != nil {
		originalRental.UserId = *updatedRental.UserID
	}
	if updatedRental.BikeID != nil {
		originalRental.BikeId = *updatedRental.BikeID
	}
	if updatedRental.Status != nil {
		originalRental.RentalStatus = *updatedRental.Status
	}
	if updatedRental.StartTime != nil {
		originalRental.StartTime = *updatedRental.StartTime
	}
	if updatedRental.EndTime != nil {
		originalRental.EndTime = updatedRental.EndTime
	}
	if updatedRental.StartLatitude != nil {
		originalRental.StartLatitude = *updatedRental.StartLatitude
	}
	if updatedRental.StartLongitude != nil {
		originalRental.StartLongitude = *updatedRental.StartLongitude
	}
	if updatedRental.EndLatitude != nil {
		originalRental.EndLatitude = updatedRental.EndLatitude
	}
	if updatedRental.EndLongitude != nil {
		originalRental.EndLongitude = updatedRental.EndLongitude
	}
	if updatedRental.Duration != nil {
		originalRental.Duration = updatedRental.Duration
	}

	_, err = rentalRepo.Update(originalRental)
	if err != nil {
		return nil, err
	}

	return originalRental, nil
}
