package handler

import (
	"net/http"
	"personal-task-tracker/internal/task"
	"personal-task-tracker/internal/view"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type TaskHandler struct {
	http.Handler

	list      *task.List
	taskView  *view.ModelView[task.Task]
	indexView *view.IndexView
}

func NewTaskHandler(service *task.List, taskView *view.ModelView[task.Task], indexView *view.IndexView) (*TaskHandler, error) {
	r := chi.NewRouter()
	h := &TaskHandler{
		Handler:   r,
		list:      service,
		taskView:  taskView,
		indexView: indexView,
	}

	r.Get("/", h.index)

	r.Post("/tasks", h.add)
	r.Get("/tasks", h.search)
	r.Post("/tasks/sort", h.reOrder)
	r.Get("/tasks/{ID}/edit", h.edit)
	r.Post("/tasks/{ID}/toggle", h.toggle)
	r.Delete("/tasks/{ID}", h.delete)
	r.Patch("/tasks/{ID}", h.rename)

	staticHandler := http.FileServer(http.Dir("static"))

	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		staticHandler.ServeHTTP(w, r)
	})

	// r.Get("/static/*", http.HandlerFunc(
	// 	func(w http.ResponseWriter, r *http.Request) {
	// 		http.StripPrefix("/static/", staticHandler).ServeHTTP(w, r)
	// 	}))

	return h, nil
}

func (t *TaskHandler) index(w http.ResponseWriter, _ *http.Request) {
	t.indexView.Index(w, t.list.Tasks())
}

func (t *TaskHandler) add(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	t.list.Add(r.FormValue("description"))
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (t *TaskHandler) search(w http.ResponseWriter, r *http.Request) {
	searchTerm := r.URL.Query().Get("search")
	results := t.list.Search(searchTerm)
	t.taskView.List(w, results)
}
func (t *TaskHandler) reOrder(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	t.list.ReOrder(r.Form["id"])
	t.taskView.List(w, t.list.Tasks())
}
func (t *TaskHandler) edit(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "ID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	item := t.list.Get(id)
	t.taskView.Edit(w, item)
}
func (t *TaskHandler) toggle(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "ID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	t.taskView.View(w, t.list.ToggleDone(id))
}
func (t *TaskHandler) delete(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "ID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	t.list.Delete(id)
}
func (t *TaskHandler) rename(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id, err := uuid.Parse(chi.URLParam(r, "ID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	desc := r.Form["name"][0]
	t.taskView.View(w, t.list.Rename(id, desc))
}
