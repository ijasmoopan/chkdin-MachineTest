package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ijasmoopan/chkdin-MachineTest/handlers"
) 

func Route() *chi.Mux {

	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/login", handlers.Login)
	router.Get("/home/{username}", handlers.GetUser)
	router.Post("/signup", handlers.PostUser)
	router.Patch("/home/{username}/patch", handlers.PatchUser)
	router.Delete("/home/{username}/delete", handlers.DeletUser)

	return router
}