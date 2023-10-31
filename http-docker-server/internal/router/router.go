package router

import (
	"github.com/go-chi/chi/v5"
)



func SetRouter(r *chi.Mux) {

	r.Get("/task", nil)
	r.Post("/task", nil)
	r.Get("/task/{id}", nil)
	r.Put("/task/{id}", nil)
	r.Delete("/task/{id}", nil)
}
