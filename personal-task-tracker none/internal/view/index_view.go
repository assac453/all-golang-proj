package view

import (
	"html/template"
	"net/http"
	"personal-task-tracker/internal/task"
)

type IndexView struct {
	templ *template.Template
}

func NewIndexView(templ *template.Template) *IndexView {
	return &IndexView{templ: templ}
}

func (i *IndexView) Index(w http.ResponseWriter, tasks []task.Task) {
	var viewModel any = tasks
	if err := i.templ.ExecuteTemplate(w, "index", viewModel); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
