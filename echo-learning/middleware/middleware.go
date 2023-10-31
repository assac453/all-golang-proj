package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func SetupMiddleware(e *echo.Echo) {
	e.Use(middleware.CORS())
	e.Use(middleware.JWT([]byte("secret-key")))
}
