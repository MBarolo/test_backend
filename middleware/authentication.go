package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/mbarolo/test_back/models"
	"github.com/mbarolo/test_back/utils"
)

type Claims struct {
	Sub       string `json:"sub"`
	Exp       int64  `json:"exp"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// Valid implementa la interfaz jwt.Claims
func (c Claims) Valid() error {
	if c.Exp < time.Now().Unix() {
		return errors.New("token expirado")
	}
	return nil
}

const EXPIRATION_IN_DAYS = 30

var jwtKey []byte

func init() {
	// Asignamos la key del .env, si no existe se asigna una default
	key := os.Getenv("JWT_KEY")
	if key == "" {
		key = "default_key"
	}

	jwtKey = []byte(key)
}

// GenerateToken: Genera un token json segun las claims del struct
func GenerateToken(user models.User) (string, time.Time, error) {
	expirationTime := time.Now().Add(time.Hour * 24 * EXPIRATION_IN_DAYS)
	claims := &Claims{
		Sub:       fmt.Sprintf("%d", user.Id),
		Exp:       expirationTime.Unix(),
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", time.Time{}, err
	}

	return tokenString, expirationTime, nil
}

// ValidateToken: Valida un token json
func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("método de firma inválido")
		}
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("token inválido")
	}

	return claims, nil
}

// Middleware que valida el token enviado, si no es encontrado retorna 401
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		token := ""

		if header != "" {
			parts := strings.Split(header, " ")
			if len(parts) == 2 && parts[0] == "Bearer" {
				token = parts[1]
			}
		} else {
			utils.JsonResponse(w, http.StatusUnauthorized, "Token no proporcionado", nil)
			return
		}

		claims, err := ValidateToken(token)
		if err != nil {
			utils.JsonResponse(w, http.StatusUnauthorized, "Error al validar el token: "+err.Error(), nil)
			return
		}

		ctx := context.WithValue(r.Context(), "claims", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
