package controller

import (
	"encoding/json"
	"net/http"

	"github.com/mbarolo/test_back/forms"
	"github.com/mbarolo/test_back/middleware"
	"github.com/mbarolo/test_back/models"
	"github.com/mbarolo/test_back/services"
	"github.com/mbarolo/test_back/utils"
	"golang.org/x/crypto/bcrypt"
)

func comparePasswords(hash string, password string) error {
	byteHash := []byte(hash)
	byteLogin := []byte(password)
	err := bcrypt.CompareHashAndPassword(byteHash, byteLogin)
	if err != nil {
		return err
	}

	return nil
}

func Login(w http.ResponseWriter, r *http.Request) {
	var loginData models.Login
	if err := json.NewDecoder(r.Body).Decode(&loginData); err != nil {
		utils.JsonResponse(w, http.StatusBadRequest, "Error al decodificar cuerpo de la solicitud: "+err.Error(), nil)
		return
	}

	if loginData.Email == "" || loginData.Password == "" {
		utils.JsonResponse(w, http.StatusBadRequest, "Email o contraseña vacíos", nil)
		return
	}

	user, err := services.GetUserByEmail(loginData.Email)
	if err != nil {
		utils.JsonResponse(w, http.StatusUnauthorized, "La información de inicio de sesión es incorrecta", nil)
		return
	}

	if err = comparePasswords(user.HashedPassword, loginData.Password); err != nil {
		utils.JsonResponse(w, http.StatusUnauthorized, "La información de inicio de sesión es incorrecta", nil)
		return
	}

	token, exp, err := middleware.GenerateToken(*user)
	if err != nil {
		utils.JsonResponse(w, http.StatusUnauthorized, "Error al iniciar sesión: ", err.Error())
	}

	utils.JsonResponse(w, http.StatusOK, "Sesión iniciada correctamente", models.LoginResponse{Token: token, Expire: exp.String()})
}

func Register(w http.ResponseWriter, r *http.Request) {
	var userForm *forms.UserForm
	if err := json.NewDecoder(r.Body).Decode(&userForm); err != nil {
		utils.JsonResponse(w, http.StatusBadRequest, "Error al decodificar el cuerpo de la solicitud: "+err.Error(), nil)
		return
	}
	defer r.Body.Close()

	// Se hashea la contraseña del usuario
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userForm.HashedPassword), bcrypt.DefaultCost)
	if err != nil {
		utils.JsonResponse(w, http.StatusInternalServerError, "Error al hashear la contraseña: "+err.Error(), nil)
		return
	}
	userForm.HashedPassword = string(hashedPassword)

	user, err := services.CreateUser(userForm.ToUser())
	if err != nil {
		utils.JsonResponse(w, http.StatusInternalServerError, "Error al crear el usuario: "+err.Error(), nil)
		return
	}

	utils.JsonResponse(w, http.StatusCreated, "Usuario creado correctamente", user)
}
