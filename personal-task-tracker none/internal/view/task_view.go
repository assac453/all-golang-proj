package view

import (
	"html/template"
	"personal-task-tracker/internal/task"
)

func NewTaskView(templ *template.Template)  *ModelView[task.Task]{
	return NewModelView[task.Task](templ, "task")
}