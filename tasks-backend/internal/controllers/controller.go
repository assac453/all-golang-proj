package controllers

import (
	"fmt"

	"github.com/assac453/tasks-backend/internal/entity"
	"github.com/assac453/tasks-backend/internal/service"
	"github.com/labstack/echo"
)

type IController interface {
	Save(echo.Context) error
	GetAll() []entity.Task
	GetById(echo.Context) (entity.Task, error)
	DeleteById(echo.Context) error
	Update(echo.Context) (entity.Task, error)
	DeleteAll() error
}

type controller struct {
	s service.IService
}

func New(s service.IService) *controller {
	return &controller{
		s: s,
	}
}

func (c *controller) Save(ctx echo.Context) error {
	var t entity.Task
	ctx.Bind(&t)
	if _, err := c.s.Save(t); err != nil {
		return err
	}
	return nil
}
func (c *controller) GetAll() []entity.Task {
	return c.s.GetAll()
}
func (c *controller) GetById(ctx echo.Context) (entity.Task, error) {
	t := struct {
		ID int `json:"id"`
	}{}
	ctx.Bind(&t)
	result, err := c.s.GetById(t.ID)
	if err != nil {
		return entity.Task{}, err
	}
	return result, nil
}
func (c *controller) DeleteById(ctx echo.Context) error {
	t := struct {
		ID int `json:"id"`
	}{}
	ctx.Bind(&t)

	return c.s.DeleteById(t.ID)
}
func (c *controller) Update(ctx echo.Context) (entity.Task, error) {
	t := entity.Task{}
	ctx.Bind(&t)

	return c.s.Update(t.ID, t)
}
func (c *controller) DeleteAll() error {
	if ok := c.s.DeleteAll(); ok {
		return nil
	}
	return fmt.Errorf("can't delete all for some reason")
}
