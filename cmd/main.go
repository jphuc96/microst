package main

import (
	"github.com/jphuc96/microst/pkg/api"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	api.Init(e)
	api.Start(e)
}
