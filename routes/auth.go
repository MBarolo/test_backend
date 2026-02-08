package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/mbarolo/test_back/controller"
)

func InitAuthRoutes(r chi.Router) {
	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", controller.Login)
		r.Post("/register", controller.Register)
	})
}
