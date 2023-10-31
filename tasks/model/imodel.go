package model

type IModel interface {
	GetAll() ([]Task, error)
	Get(id int) (Task, error)
	Delete(id int) error
	Update(id int, new Task) error
	Save(t Task) error
	Count() int
}
