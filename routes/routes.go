package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mbarolo/test_back/utils"
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

		r.Get("/status", func(w http.ResponseWriter, r *http.Request) {
			utils.JsonResponse(w, http.StatusOK, "ok", nil)
		})
	})
}
