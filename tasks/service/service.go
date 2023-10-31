package service

import (
	"github.com/assac453/tasks/model"
)

type IService interface {
	GetAll() ([]model.Task, error)
	Get(id int) (model.Task, error)
	Delete(id int) error
	Update(id int, t model.Task) error
	Save(t model.Task) error
}

type service struct {
	model model.IModel
}

func New(model model.IModel) IService {
	return &service{
		model: model,
	}
}

func (s *service) GetAll() ([]model.Task, error) {
	return s.model.GetAll()
}
func (s *service) Get(id int) (model.Task, error) {
	return s.model.Get(id)
}
func (s *service) Delete(id int) error {
	return s.model.Delete(id)
}
func (s *service) Update(id int, t model.Task) error {
	return s.model.Update(id, t)
}
func (s *service) Save(t model.Task) error {
	return s.model.Save(t)
}
