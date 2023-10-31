package main

import (
	"log"
	"net/http"
	handler "personal-task-tracker/internal"
	"personal-task-tracker/internal/task"
	"personal-task-tracker/internal/view"
)

const addr = ":8080"

func main() {
	service := task.List{}
	service.Add("some1")
	service.Add("some2")
	service.Add("some3")

	templates, err := view.NewTemplates()
	if err != nil {
		log.Fatal(err)
	}

	h, err := handler.NewTaskHandler(&service, view.NewTaskView(templates), view.NewIndexView(templates))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("listening on : %v", addr)

	log.Fatal(http.ListenAndServe(addr, h))
}
