package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/mbarolo/test_back/controller"
)

func InitAdminRoutes(r chi.Router) {
	r.Route("/admin", func(r chi.Router) {
		//todo middleware para admin
		r.Post("/bikes", controller.CreateBike)
		r.Patch("/bikes/{id}", controller.UpdateBike)
		r.Get("/bikes", controller.GetAllBikes)
	})
}
