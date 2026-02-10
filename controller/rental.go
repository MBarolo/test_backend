package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mbarolo/test_back/forms"
	"github.com/mbarolo/test_back/services"
	"github.com/mbarolo/test_back/utils"
)

// StartRental godoc
// @Summary      Iniciar alquiler
// @Description  Alquilar una bicicleta disponible
// @Tags         rentals
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        rental  body      forms.StartEndRentalForm  true  "ID de la bicicleta a alquilar"
// @Success      201     {object}  map[string]interface{}
// @Failure      400     {object}  map[string]interface{}
// @Failure      401     {object}  map[string]interface{}
// @Failure      500     {object}  map[string]interface{}
// @Router       /rentals/start [post]
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

// EndRental godoc
// @Summary      Finalizar alquiler
// @Description  Finalizar el alquiler activo de una bicicleta
// @Tags         rentals
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        rental  body      forms.StartEndRentalForm  true  "ID de la bicicleta a devolver"
// @Success      200     {object}  map[string]interface{}
// @Failure      400     {object}  map[string]interface{}
// @Failure      401     {object}  map[string]interface{}
// @Failure      500     {object}  map[string]interface{}
// @Router       /rentals/end [post]
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

// GetUserRentalHistory godoc
// @Summary      Historial de alquileres
// @Description  Obtener el historial de alquileres del usuario autenticado
// @Tags         rentals
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  map[string]interface{}
// @Failure      401  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /rentals/history [get]
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

// GetAllRentals godoc
// @Summary      Obtener todos los alquileres
// @Description  Listar todos los alquileres del sistema (admin)
// @Tags         admin
// @Accept       json
// @Produce      json
// @Security     BasicAuth
// @Success      200  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /admin/rentals [get]
func GetAllRentals(w http.ResponseWriter, r *http.Request) {
	rentals, err := services.GetAllRentals()
	if err != nil {
		utils.JsonResponse(w, http.StatusInternalServerError, "Error al obtener los alquileres: "+err.Error(), nil)
		return
	}

	utils.JsonResponse(w, http.StatusOK, "Alquileres obtenidos", rentals)
}

// GetRentalById godoc
// @Summary      Obtener alquiler por ID
// @Description  Obtener información de un alquiler específico (admin)
// @Tags         admin
// @Accept       json
// @Produce      json
// @Security     BasicAuth
// @Param        id   path      int  true  "ID del alquiler"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /admin/rentals/{id} [get]
func GetRentalById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		utils.JsonResponse(w, http.StatusBadRequest, "Parametro id no encontrado", nil)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.JsonResponse(w, http.StatusInternalServerError, "Error al convertir el parametro id: "+err.Error(), nil)
		return
	}

	rental, err := services.GetRentalById(int64(id))
	if err != nil {
		utils.JsonResponse(w, http.StatusInternalServerError, "Error al obtener el alquiler: "+err.Error(), nil)
		return
	}

	utils.JsonResponse(w, http.StatusOK, "Alquiler obtenido", rental)
}

// UpdateRental godoc
// @Summary      Actualizar alquiler
// @Description  Modificar información de un alquiler específico (admin)
// @Tags         admin
// @Accept       json
// @Produce      json
// @Security     BasicAuth
// @Param        id      path      int              true  "ID del alquiler"
// @Param        rental  body      forms.RentalForm true  "Datos actualizados del alquiler"
// @Success      200     {object}  map[string]interface{}
// @Failure      400     {object}  map[string]interface{}
// @Failure      500     {object}  map[string]interface{}
// @Router       /admin/rentals/{id} [patch]
func UpdateRental(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		utils.JsonResponse(w, http.StatusBadRequest, "Parametro id no encontrado", nil)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.JsonResponse(w, http.StatusInternalServerError, "Error al convertir el parametro id: "+err.Error(), nil)
		return
	}

	var rentalForm *forms.RentalForm
	if err := json.NewDecoder(r.Body).Decode(&rentalForm); err != nil {
		utils.JsonResponse(w, http.StatusBadRequest, "Error al decodificar el cuerpo de la solicitud: "+err.Error(), nil)
		return
	}

	rental, err := services.UpdateRental(int64(id), rentalForm)
	if err != nil {
		utils.JsonResponse(w, http.StatusInternalServerError, "Error al actualizar el alquiler: "+err.Error(), nil)
		return
	}

	utils.JsonResponse(w, http.StatusOK, "Alquiler actualizado", rental)
}
