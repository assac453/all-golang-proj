package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	count := 0
	r.Handle("/", http.FileServer(http.Dir("static")))

	r.Get("/about", func(w http.ResponseWriter, r *http.Request) {
		tmplt := template.Must(template.ParseFiles("static/about.html"))
		tmplt.Execute(w, nil)
	})
	r.Post("/clicked", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte(fmt.Sprintf("%v clicked", count)))
	})
	s := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	log.Fatal(s.ListenAndServe())
}
