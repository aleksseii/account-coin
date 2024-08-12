package handler

import (
	"account-coin/internal/middleware"
	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
)

func Handler(router *chi.Mux) {
	router.Use(chimiddleware.StripSlashes)
	router.Route("/accounts", func(r chi.Router) {

		r.Use(middleware.Authorization)

		r.Get("/coins", GetCoinBalance)
	})
}
