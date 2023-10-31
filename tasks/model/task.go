package model

import (
	"fmt"

	"gorm.io/gorm"
)

type Task_gorm struct {
	gorm.Model
	Title     string `json:"title"`
	Completed string `json:"completed"`
}

func (t Task_gorm) String() string {
	return fmt.Sprintf("%v, %s, %s", t.ID, t.Title, t.Completed)
}

type Task_sqlite struct {
	ID        string
	Title     string `json:"title"`
	Completed string `json:"completed"`
}

func (t Task_sqlite) String() string {
	return fmt.Sprintf("%v, %v, %v", t.ID, t.Title, t.Completed)
}
