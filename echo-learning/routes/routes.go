package routes

import (
	"github.com/assac453/echo-learning/app/handlers"
	"github.com/labstack/echo"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/user", handlers.GetUserByEmail)

	e.GET("/users", handlers.GetAllUsers)

	e.POST("/user", handlers.AddUserWithotTasks)

	e.POST("/user/todo", handlers.AddTaskForUser)

	e.POST("/login", handlers.Login)
}
