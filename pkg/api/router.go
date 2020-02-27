package api

import (
	"github.com/jphuc96/microst/pkg/api/system"
	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	system.Init(e)
}

func Start(e *echo.Echo) {
	e.Logger.Fatal(e.Start("0.0.0.0:8080"))
}
