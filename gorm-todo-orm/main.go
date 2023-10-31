package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	setRoutes(r)

	s := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	s.ListenAndServe()
}

func setRoutes(r *chi.Mux) {
	r.Get("/{}", func(w http.ResponseWriter, r *http.Request) {

	})
}
