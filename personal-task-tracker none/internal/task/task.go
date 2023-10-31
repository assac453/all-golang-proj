package task

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

type Task struct {
	ID          uuid.UUID
	Description string
	CreatedAt   time.Time
	Complete    bool
}

type List struct {
	tasks []Task
}

// add new Task with description
func (l *List) Add(description string) {
	l.tasks = append(l.tasks, Task{
		ID:          uuid.New(),
		Description: description,
		CreatedAt:   time.Now(),
	})
}

// rename description by id, new description
func (l *List) Rename(id uuid.UUID, description string) Task {
	i := l.indexOf(id)
	l.tasks[i].Description = description
	return l.tasks[i]
}

func (l *List) Tasks() []Task {
	return l.tasks
}

// toggle completed flag by id
func (l *List) ToggleDone(id uuid.UUID) Task {
	i := l.indexOf(id)
	l.tasks[i].Complete = !l.tasks[i].Complete
	return l.tasks[i]
}

// delete by id
func (l *List) Delete(id uuid.UUID) {
	i := l.indexOf(id)
	l.tasks = append(l.tasks[:i], l.tasks[i+1:]...)
}

// reoder in slice by uuid
func (l *List) ReOrder(ids []string) {
	var uuids []uuid.UUID
	for _, id := range ids {
		uuids = append(uuids, uuid.MustParse(id))
	}

	var newList []Task
	for _, id := range uuids {
		newList = append(newList, l.tasks[l.indexOf(id)])
	}

	l.tasks = newList
}

// return slice of result by str description
func (l *List) Search(search string) []Task {
	search = strings.ToLower(search)
	var results []Task
	for _, v := range l.tasks {
		if strings.Contains(strings.ToLower(v.Description), search) {
			results = append(results, v)
		}
	}
	return results
}

// return Task by uuid.UUID
func (l *List) Get(id uuid.UUID) Task {
	return l.tasks[l.indexOf(id)]
}

// make current List empty
func (l *List) Empty() {
	l.tasks = []Task{}
}

// return index of item by uuid.UUID
func (l *List) indexOf(id uuid.UUID) int {
	return slices.IndexFunc(l.tasks, func(t Task) bool {
		return t.ID == id
	})
}
