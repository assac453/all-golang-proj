package model

import (
	"fmt"

	"github.com/assac453/tasks-backend/internal/entity"
)

var Tasks map[int]entity.Task

func init() {
	Tasks = map[int]entity.Task{
		1: {
			ID:          1,
			Title:       "Some title",
			Description: "Do some things and be handsome",
			Status:      "in progress",
		},
	}
}

func Save(t entity.Task) error {
	if _, ok := Tasks[t.ID]; ok {
		return nil
	}
	Tasks[t.ID] = t
	return nil
}
func GetAll() []entity.Task {
	var result []entity.Task
	for _, v := range Tasks {
		result = append(result, v)
	}
	return result
}
func GetById(id int) (entity.Task, error) {
	if _, ok := Tasks[id]; !ok {
		return entity.Task{}, fmt.Errorf("task with id:%d not found", id)
	}
	return Tasks[id], nil
}
func DeleteById(id int) error {
	if _, ok := Tasks[id]; !ok {
		return fmt.Errorf("task with id:%d not found", id)
	}
	delete(Tasks, id)
	return nil
}
func Update(id int, t entity.Task) error {
	if _, ok := Tasks[id]; !ok {
		return fmt.Errorf("task with id:%d not found", id)
	}
	Tasks[id] = t
	return nil
}
func DeleteAll() error {
	Tasks = make(map[int]entity.Task)
	return nil
}
