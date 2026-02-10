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

// GetAvailableBikes godoc
// @Summary      Obtener bicicletas disponibles
// @Description  Retorna las bicicletas disponibles para alquilar
// @Tags         bikes
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /bikes/available [get]
func GetAvailableBikes(w http.ResponseWriter, r *http.Request) {
	bikes, err := services.GetAvailableBikes()
	if err != nil {
		utils.JsonResponse(w, http.StatusInternalServerError, "Error al obtener las bicicletas:"+err.Error(), nil)
		return
	}

	utils.JsonResponse(w, http.StatusOK, "Bicicletas obtenidas", bikes)
}

// GetAllBikes godoc
// @Summary      Obtener todas las bicicletas
// @Description  Retorna todas las bicicletas del sistema (admin)
// @Tags         admin
// @Accept       json
// @Produce      json
// @Security     BasicAuth
// @Success      200  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /admin/bikes [get]
func GetAllBikes(w http.ResponseWriter, r *http.Request) {
	bikes, err := services.GetAllBikes()
	if err != nil {
		utils.JsonResponse(w, http.StatusInternalServerError, "Error al obtener las bicicletas: "+err.Error(), nil)
		return
	}

	utils.JsonResponse(w, http.StatusOK, "Bicicletas obtenidas", bikes)
}

// CreateBike godoc
// @Summary      Crear bicicleta
// @Description  Registrar una nueva bicicleta en el sistema (admin)
// @Tags         admin
// @Accept       json
// @Produce      json
// @Security     BasicAuth
// @Param        bike  body      forms.BikeForm  true  "Datos de la nueva bicicleta"
// @Success      201   {object}  map[string]interface{}
// @Failure      400   {object}  map[string]interface{}
// @Failure      500   {object}  map[string]interface{}
// @Router       /admin/bikes [post]
func CreateBike(w http.ResponseWriter, r *http.Request) {
	var bikeForm *forms.BikeForm
	if err := json.NewDecoder(r.Body).Decode(&bikeForm); err != nil {
		utils.JsonResponse(w, http.StatusBadRequest, "Error al decodificar el cuerpo de la solicitud: "+err.Error(), nil)
		return
	}

	bike, err := services.CreateBike(bikeForm.ToBike())
	if err != nil {
		utils.JsonResponse(w, http.StatusInternalServerError, "Error al crear la bicicleta: "+err.Error(), nil)
		return
	}

	utils.JsonResponse(w, http.StatusCreated, "Bicicleta creada correctamente", bike)
}

// UpdateBike godoc
// @Summary      Actualizar bicicleta
// @Description  Modificar los datos de una bicicleta existente (admin)
// @Tags         admin
// @Accept       json
// @Produce      json
// @Security     BasicAuth
// @Param        id    path      int             true  "ID de la bicicleta"
// @Param        bike  body      forms.BikeForm  true  "Datos actualizados de la bicicleta"
// @Success      200   {object}  map[string]interface{}
// @Failure      400   {object}  map[string]interface{}
// @Failure      500   {object}  map[string]interface{}
// @Router       /admin/bikes/{id} [patch]
func UpdateBike(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		utils.JsonResponse(w, http.StatusBadRequest, "Parametro id no encontrado", nil)
		return
	}

	var bikeForm *forms.BikeForm
	if err := json.NewDecoder(r.Body).Decode(&bikeForm); err != nil {
		utils.JsonResponse(w, http.StatusBadRequest, "Error al decodificar el cuerpo de la solicitud: "+err.Error(), nil)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.JsonResponse(w, http.StatusInternalServerError, "Error al convertir el parametro id: "+err.Error(), nil)
		return
	}

	bike, err := services.UpdateBike(int64(id), bikeForm)
	if err != nil {
		utils.JsonResponse(w, http.StatusInternalServerError, "Error al actualizar la bicicleta: "+err.Error(), nil)
		return
	}

	utils.JsonResponse(w, http.StatusOK, "Bicicleta actualizda correctamente", bike)
}
