package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// checkVars: Función que revisa si falta alguna variable de entorno necesaria en el archivo .env
func checkVars() []string {
	vars := []string{"ADDR", "JWT_SECRET", "ADMIN_CREDENTIALS"}
	missing := []string{}
	for _, v := range vars {
		_, set := os.LookupEnv(v)
		if !set {
			missing = append(missing, v)
		}
	}
	return missing
}

// LoadEnv: Función que carga las variables de entorno
func LoadEnv() {
	godotenv.Load(".env")
	if vars := checkVars(); len(vars) != 0 {
		log.Printf("ERROR: Variables de entorno necesarias no definidas: %v", vars)
		panic(fmt.Sprintf("ERROR: Variables de entorno necesarias no definidas: %v", vars))
	}
}
