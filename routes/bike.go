package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/mbarolo/test_back/controller"
	"github.com/mbarolo/test_back/middleware"
)

func InitBikeRoutes(r chi.Router) {
	r.Route("/bikes", func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)
		r.Get("/available", controller.GetAvailableBikes)
	})
}
