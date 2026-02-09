package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/mbarolo/test_back/controller"
	"github.com/mbarolo/test_back/middleware"
)

func InitRentalRoutes(r chi.Router) {
	r.Route("/rentals", func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)
		r.Get("/history", controller.GetUserRentalHistory)
		r.Post("/start", controller.StartRental)
		r.Post("/end", controller.EndRental)
	})
}
