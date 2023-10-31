package view

import (
	"html/template"
	"net/http"
)

type ModelView[T any] struct {
	templ     *template.Template
	modelName string
}

func NewModelView[T any](templ *template.Template, modelName string) *ModelView[T] {
	return &ModelView[T]{templ: templ, modelName: modelName}
}

func (m *ModelView[T]) List(w http.ResponseWriter, results []T) {
	m.renderOr500(w, m.modelName+"s", results)
}

func (m *ModelView[T]) View(w http.ResponseWriter, item T) {
	m.renderOr500(w, "view_"+m.modelName, item)
}
func (m *ModelView[T]) Edit(w http.ResponseWriter, item T) {
	m.renderOr500(w, "edit_"+m.modelName, item)
}
func (m *ModelView[T]) renderOr500(w http.ResponseWriter, templateName string, viewModel any) {
	if err := m.templ.ExecuteTemplate(w, templateName, viewModel); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
