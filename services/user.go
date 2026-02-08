package services

import (
	"errors"
	"log"
	"time"

	"github.com/mbarolo/test_back/forms"
	"github.com/mbarolo/test_back/models"
)

func GetAllUsers() ([]*models.User, error) {
	if users, err := userRepo.GetAll(); err != nil {
		log.Printf("Error al obtener los usuarios: %v", err.Error())
		return nil, err
	} else {
		log.Println("Usuarios obtenidos")
		return users, nil
	}
}

func GetUserById(id int64) (*models.User, error) {
	if user, err := userRepo.GetById(id); err != nil {
		log.Printf("Error al obtener el usuario: %v", err.Error())
		return nil, err
	} else {
		log.Println("Usuario obtenido")
		return user, nil
	}
}

func GetUserByEmail(email string) (*models.User, error) {
	if user, err := userRepo.GetByEmail(email); err != nil {
		log.Printf("Error al obtener el usuario: %v", err.Error())
		return nil, err
	} else {
		log.Println("Usuario obtenido")
		return user, nil
	}
}

func CreateUser(user *models.User) (*models.User, error) {
	if err := user.ValidateFields(); err != nil {
		return nil, err
	}

	exists, err := userRepo.ExistsByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, errors.New("correo electrónico ya registrado en la base de datos")

	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	id, err := userRepo.Create(user)
	if err != nil {
		log.Printf("Error al crear el usuario: %v", err.Error())
		return nil, err
	}

	user.Id = id

	log.Println("Usuario creado exitosamente")
	return user, nil
}

func UpdateUser(id int64, updatedUser *forms.UserForm) (*models.User, error) {
	originalUser, err := GetUserById(id)
	if err != nil {
		return nil, err
	}

	if updatedUser.Email != "" {
		originalUser.Email = updatedUser.Email
	}
	if updatedUser.FirstName != "" {
		originalUser.FirstName = updatedUser.FirstName
	}
	if updatedUser.LastName != "" {
		originalUser.LastName = updatedUser.LastName
	}
	if updatedUser.HashedPassword != "" {
		originalUser.HashedPassword = updatedUser.HashedPassword
	}

	if err := originalUser.ValidateFields(); err != nil {
		return nil, err
	}

	originalUser.UpdatedAt = time.Now()

	_, err = userRepo.Update(originalUser)
	if err != nil {
		return nil, err
	}

	return originalUser, nil
}

func DeleteUser(id int64) error {
	originalUser, err := GetUserById(id)
	if err != nil {
		return err
	}

	originalUser.Deleted = true

	if _, err := userRepo.Update(originalUser); err != nil {
		return err
	}

	return nil
}
