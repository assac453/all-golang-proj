package main

import (
	"log"

	"github.com/assac453/tasks-backend/internal/middleware"
	"github.com/assac453/tasks-backend/internal/routes"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	middleware.SetupMiddleware(e)

	routes.SetupRouter(e)

	log.Fatal(e.Start(":8080"))
}
