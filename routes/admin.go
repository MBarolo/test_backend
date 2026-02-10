package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/mbarolo/test_back/controller"
	"github.com/mbarolo/test_back/middleware"
)

func InitAdminRoutes(r chi.Router) {
	r.Route("/admin", func(r chi.Router) {
		r.Use(middleware.AdminMiddleware)

		r.Post("/bikes", controller.CreateBike)
		r.Patch("/bikes/{id}", controller.UpdateBike)
		r.Get("/bikes", controller.GetAllBikes)

		r.Get("/users", controller.GetAllUsers)
		r.Get("/users/{id}", controller.GetUserById)
		r.Patch("/users/{id}", controller.UpdateUser)

		r.Get("/rentals", controller.GetAllRentals)
		r.Get("/rentals/{id}", controller.GetRentalById)
		r.Patch("/rentals/{id}", controller.UpdateRental)
	})
}
