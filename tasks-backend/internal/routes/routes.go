package routes

import (
	"net/http"

	"github.com/assac453/tasks-backend/internal/controllers"
	"github.com/assac453/tasks-backend/internal/service"
	"github.com/labstack/echo"
)

var (
	s service.IService        = service.New()
	c controllers.IController = controllers.New(s)
)

func SetupRouter(e *echo.Echo) {

	e.GET("/tasks", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, c.GetAll())
	})
}
