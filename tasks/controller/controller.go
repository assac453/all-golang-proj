package controller

import (
	"github.com/assac453/tasks/model"
	"github.com/assac453/tasks/service"
	"github.com/labstack/echo"
	"github.com/rs/zerolog/log"
)

type IController interface {
	GetAll() []model.Task
	Get(id int) model.Task
	Delete(id int)
	Update(id int, t model.Task)
	Save(ctx echo.Context)
}

type controller struct {
	service service.IService
}

func New(service service.IService) IController {
	return &controller{
		service: service,
	}
}

func (c *controller) GetAll() []model.Task {
	tasks, err := c.service.GetAll()
	if err != nil {
		return []model.Task{}
	}
	return tasks
}

func (c *controller) Get(id int) model.Task {
	task, err := c.service.Get(id)
	if err != nil {
		return nil
	}
	return task
}

func (c *controller) Delete(id int) {
	err := c.service.Delete(id)
	if err != nil {
		log.Error().Msg(err.Error())
	}
}

func (c *controller) Update(id int, t model.Task) {
	err := c.service.Update(id, t)
	if err != nil {
		log.Error().Msg(err.Error())
	}
}

func (c *controller) Save(ctx echo.Context) {
	var t model.Task
	err := ctx.Bind(&t)
	if err != nil {
		log.Error().Msg(err.Error())
	}
	err = c.service.Save(t)
	if err != nil {
		log.Error().Msg(err.Error())
	}
}
