package services

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/mbarolo/test_back/middleware"
	"github.com/mbarolo/test_back/models"
)

func GetCurrentUser(r *http.Request) (*models.User, error) {
	// Se obtienen los claims del token
	claims, ok := r.Context().Value("claims").(*middleware.Claims)
	if !ok {
		return nil, errors.New("Error al validar los claims del token")
	}

	// Sacamos el id del user logeado
	userID, err := strconv.ParseInt(claims.Sub, 10, 64)
	if err != nil {
		return nil, err
	}

	user, err := GetUserById(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
