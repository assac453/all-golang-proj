package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/assac453/tasks/controller"
	"github.com/assac453/tasks/model"
	model_sqlite "github.com/assac453/tasks/model/sqlite_driver"
	"github.com/assac453/tasks/service"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	m model.IModel           = model_sqlite.New()
	s service.IService       = service.New(m)
	c controller.IController = controller.New(s)
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger(), middleware.CORS(), middleware.Recover())

	e.GET("/task", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, c.GetAll())
	})

	e.GET("/task/:id", func(ctx echo.Context) error {
		param_id := ctx.Param("id")
		id, _ := strconv.Atoi(param_id)
		return ctx.JSON(http.StatusOK, c.Get(id))
	})

	e.DELETE("/task/:id", func(ctx echo.Context) error {
		param_id := ctx.Param("id")
		id, _ := strconv.Atoi(param_id)
		t := c.Get(id)
		c.Delete(id)
		return ctx.JSON(http.StatusOK, t)
	})

	e.PUT("/task/:id", func(ctx echo.Context) error {
		param_id := ctx.Param("id")
		id, _ := strconv.Atoi(param_id)

		var t model.Task
		ctx.Bind(&t)

		c.Update(id, t)

		return ctx.JSON(http.StatusOK, t)
	})

	e.POST("/task", func(ctx echo.Context) error {
		var t model.Task

		ctx.Bind(&t)

		c.Save(ctx)

		return ctx.JSON(http.StatusOK, t)
	})

	log.Fatal(e.Start(":8080"))
}
