package server

import (
	"github.com/labstack/echo"
	"github.com/radio-noise-project/sisters/internal/api/handler"
)

func router(e *echo.Echo) {
	e.GET("/v0/runtime/version", handler.OutputSistersVersion)
}
