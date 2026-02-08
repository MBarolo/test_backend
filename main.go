package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/mbarolo/test_back/config"
	"github.com/mbarolo/test_back/routes"
	"github.com/mbarolo/test_back/utils"
)

func main() {

	// se cargan las variables de entorno
	utils.LoadEnv()

	// logging
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("Start test_back")
	log.Printf("serverUp, %s", os.Getenv("ADDR"))

	defer config.CloseDB()

	// se configura go-chi
	app := chi.NewRouter()
	app.Use(chimiddleware.Logger)
	app.Use(chimiddleware.Recoverer)

	// se define la response default para 404
	app.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": "Servicio no encontrado."})
	})

	// registramos las rutas en la aplicación
	routes.InitRoutes(app)
	chi.Walk(app, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("[%s] %s", method, route)
		return nil
	})

	log.Printf("Server starting on %s", os.Getenv("ADDR"))
	http.ListenAndServe(os.Getenv("ADDR"), app)
}
