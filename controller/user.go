package controller

import (
	"encoding/json"
	"net/http"

	"github.com/mbarolo/test_back/forms"
	"github.com/mbarolo/test_back/services"
	"github.com/mbarolo/test_back/utils"
)

func GetProfile(w http.ResponseWriter, r *http.Request) {
	user, err := services.GetCurrentUser(r)
	if err != nil {
		utils.JsonResponse(w, http.StatusUnauthorized, "Error al obtener el usuario autenticado: "+err.Error(), nil)
		return
	}

	utils.JsonResponse(w, http.StatusOK, "Perfil obtenido", user)
}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	user, err := services.GetCurrentUser(r)
	if err != nil {
		utils.JsonResponse(w, http.StatusUnauthorized, "Error al obtener el usuario autenticado: "+err.Error(), nil)
		return
	}

	var userForm *forms.UserForm
	if err := json.NewDecoder(r.Body).Decode(&userForm); err != nil {
		utils.JsonResponse(w, http.StatusBadRequest, "Error al decodificar el cuerpo de la solicitud: "+err.Error(), nil)
		return
	}

	updatedUser, err := services.UpdateUser(user.Id, userForm)
	if err != nil {
		utils.JsonResponse(w, http.StatusInternalServerError, "Error al actualizar el usuario: "+err.Error(), nil)
		return
	}

	utils.JsonResponse(w, http.StatusOK, "Usuario actualizado correctamente", updatedUser)
}
