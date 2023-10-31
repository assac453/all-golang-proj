package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type ViewData struct {
	Title   string
	Message string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := ViewData{
			Title:   "Hello world",
			Message: "My niggass",
		}
		tmpl, _ := template.ParseFiles("index.html")
		tmpl.Execute(w, data)
	})
	fmt.Println("Server is listening")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
