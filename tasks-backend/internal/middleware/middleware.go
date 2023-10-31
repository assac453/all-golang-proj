package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func SetupMiddleware(c *echo.Echo) {
	c.Use(middleware.CORS())
}
