package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Rotuer struct {
	mux *http.ServeMux
}

func InitRoutes(r chi.Router) {
	r.Route("/api/v1", func(r chi.Router) {
		InitAuthRoutes(r)
		InitUserRoutes(r)
		InitBikeRoutes(r)
		InitRentalRoutes(r)
		InitAdminRoutes(r)
	})
}
