package view

import (
	"embed"
	"html/template"
)

var (
	//go:embed "templates/*"
	taskTemplates embed.FS
)

func NewTemplates() (*template.Template, error){
	return template.ParseFS(taskTemplates, "templates/*/*.gohtml", "templates/*.gohtml")	
}