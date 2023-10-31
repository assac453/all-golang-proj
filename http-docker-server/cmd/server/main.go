package main

import (
	"net/http"
	"time"

	"github.com/assac453/http-docker-server/internal/router"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	router.SetRouter(r)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	s.ListenAndServe()
}
