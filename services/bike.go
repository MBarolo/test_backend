package services

import (
	"log"
	"time"

	"github.com/mbarolo/test_back/forms"
	"github.com/mbarolo/test_back/models"
)

func GetAvailableBikes() ([]*models.Bike, error) {
	if bikes, err := bikeRepo.GetAllAvailable(); err != nil {
		log.Printf("Error al obtener las bicicletas: %v", err.Error())
		return nil, err
	} else {
		log.Printf("Bicicletas obtenidas")
		return bikes, nil
	}
}

func GetAllBikes() ([]*models.Bike, error) {
	if bikes, err := bikeRepo.GetAll(); err != nil {
		log.Printf("Error al obtener las bicicletas: %v", err.Error())
		return nil, err
	} else {
		log.Printf("Bicicletas obtenidas")
		return bikes, nil
	}
}

func GetBikeById(id int64) (*models.Bike, error) {
	if bike, err := bikeRepo.GetById(id); err != nil {
		log.Printf("Error al obtener la bicicleta: %v", err.Error())
		return nil, err
	} else {
		log.Printf("Bicicletas obtenidas")
		return bike, nil
	}
}

func CreateBike(bike *models.Bike) (*models.Bike, error) {
	if err := bike.ValidateFields(); err != nil {
		return nil, err
	}

	bike.CreatedAt = time.Now()
	bike.UpdatedAt = time.Now()
	bike.IsAvailable = true

	id, err := bikeRepo.CreateBike(bike)
	if err != nil {
		return nil, err
	}

	bike.Id = id

	return bike, nil
}

func UpdateBike(id int64, updatedBike *forms.BikeForm) (*models.Bike, error) {
	originalBike, err := GetBikeById(id)
	if err != nil {
		return nil, err
	}

	if updatedBike.CostPerMinute != nil {
		originalBike.CostPerMinute = *updatedBike.CostPerMinute
	}
	if updatedBike.IsAvailable != nil {
		originalBike.IsAvailable = *updatedBike.IsAvailable
	}
	if updatedBike.Longitude != nil {
		originalBike.Longitude = *updatedBike.Longitude
	}
	if updatedBike.Latitude != nil {
		originalBike.Latitude = *updatedBike.Latitude
	}

	if err := originalBike.ValidateFields(); err != nil {
		return nil, err
	}

	originalBike.UpdatedAt = time.Now()

	_, err = bikeRepo.UpdateBike(originalBike)
	if err != nil {
		return nil, err
	}

	return originalBike, nil
}
