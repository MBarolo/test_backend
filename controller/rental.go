package controller

import (
	"encoding/json"
	"net/http"

	"github.com/mbarolo/test_back/forms"
	"github.com/mbarolo/test_back/services"
	"github.com/mbarolo/test_back/utils"
)

func StartRental(w http.ResponseWriter, r *http.Request) {
	user, err := services.GetCurrentUser(r)
	if err != nil {
		utils.JsonResponse(w, http.StatusUnauthorized, "Error al obtener el usuario autenticado: "+err.Error(), nil)
		return
	}

	var rentalForm *forms.StartEndRentalForm
	if err := json.NewDecoder(r.Body).Decode(&rentalForm); err != nil {
		utils.JsonResponse(w, http.StatusBadRequest, "Error al decodificar el cuerpo de la solicitud: "+err.Error(), nil)
		return
	}

	rental, err := services.StartRental(user, rentalForm)
	if err != nil {
		utils.JsonResponse(w, http.StatusInternalServerError, "Error al alquilar la bicicleta: "+err.Error(), nil)
		return
	}

	utils.JsonResponse(w, http.StatusCreated, "Bicicleta alquilada correctamente", rental)
}

func EndRental(w http.ResponseWriter, r *http.Request) {
	user, err := services.GetCurrentUser(r)
	if err != nil {
		utils.JsonResponse(w, http.StatusUnauthorized, "Error al obtener el usuario autenticado: "+err.Error(), nil)
		return
	}

	var rentalForm *forms.StartEndRentalForm
	if err := json.NewDecoder(r.Body).Decode(&rentalForm); err != nil {
		utils.JsonResponse(w, http.StatusBadRequest, "Error al decodificar el cuerpo de la solicitud: "+err.Error(), nil)
		return
	}

	rental, err := services.EndRental(user, rentalForm)
	if err != nil {
		utils.JsonResponse(w, http.StatusInternalServerError, "Error al finalizar el alquiler: "+err.Error(), nil)
		return
	}

	utils.JsonResponse(w, http.StatusOK, "Alquiler de bicicleta finalizado correctamente", rental)
}

func GetUserRentalHistory(w http.ResponseWriter, r *http.Request) {
	user, err := services.GetCurrentUser(r)
	if err != nil {
		utils.JsonResponse(w, http.StatusUnauthorized, "Error al obtener el usuario autenticado: "+err.Error(), nil)
		return
	}

	rentals, err := services.GetRentalHistory(user.Id)
	if err != nil {
		utils.JsonResponse(w, http.StatusInternalServerError, "Error al obtener el historial de alquileres: "+err.Error(), nil)
		return
	}
	utils.JsonResponse(w, http.StatusOK, "Historial de alquileres obtenido correctamente", rentals)
}
