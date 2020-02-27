package system

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/jphuc96/microst/pkg/models"
)

func Init(e *echo.Echo) {
	s := e.Group("/system")
	s.GET("/healthz", healthcheck)
}

func healthcheck(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, models.SystemStatus{
		Code:    http.StatusOK,
		Message: "ok",
	}, "  ")
}
