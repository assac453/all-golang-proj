package main

import (
	"github.com/assac453/echo-learning/middleware"
	"github.com/assac453/echo-learning/routes"

	"github.com/labstack/echo"
)

func main() {

	e := echo.New()

	middleware.SetupMiddleware(e)

	routes.SetupRoutes(e)

	e.Logger.Fatal(e.Start(":8000"))

}
