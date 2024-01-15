package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/minij147/goapi_example/internal/middleware"
)

func Handler(r *chi.Mux) {
	//global middleware
	r.Use(chimiddle.StripSlashes) //ignores trailing slashes

	r.Route("/account", func(router chi.Router) {
		router.Use(middleware.Authorization)

		router.Get("/coins", GetCoinBalance)
	})
}
