package service

import (
	"log"

	"github.com/assac453/tasks-backend/internal/entity"
	"github.com/assac453/tasks-backend/internal/model"
)

type IService interface {
	Save(entity.Task) (entity.Task, error)
	GetAll() []entity.Task
	GetById(id int) (entity.Task, error)
	DeleteById(id int) error
	Update(id int, t entity.Task) (entity.Task, error)
	DeleteAll() bool
}

type service struct{}

func New() *service {
	return &service{}
}

func (s *service) Save(t entity.Task) (entity.Task, error) {
	if err := model.Save(t); err != nil {
		log.Println(err)
		return entity.Task{}, err
	}
	return t, nil
}

func (s *service) GetAll() []entity.Task {
	return model.GetAll()
}

func (s *service) GetById(id int) (entity.Task, error) {
	result, err := model.GetById(id)
	if err != nil {
		log.Println(err)
		return entity.Task{}, err
	}
	return result, nil
}

func (s *service) DeleteById(id int) error {
	if err := model.DeleteById(id); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s *service) Update(id int, t entity.Task) (entity.Task, error) {
	if err := model.Update(id, t); err != nil {
		return entity.Task{}, err
	}
	return t, nil
}

func (s *service) DeleteAll() bool {
	if err := model.DeleteAll(); err != nil {
		return false
	}
	return true
}
