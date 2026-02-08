package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/mbarolo/test_back/controller"
	"github.com/mbarolo/test_back/middleware"
)

func InitUserRoutes(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)
		r.Get("/profile", controller.GetProfile)
		r.Patch("/profile", controller.UpdateProfile)
	})
}
