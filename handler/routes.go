package handler

import (
	"github.com/labstack/echo/v4"
)

func (h *httpHandler) initRoutes(e *echo.Echo) {
	e.GET("/health", h.HealthCheck)

	v1 := e.Group("/v1")
	v1.POST("/signup", h.SignUp)
}
