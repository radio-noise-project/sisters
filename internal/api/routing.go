package api

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/radio-noise-project/sisters/internal/api/handler"
)

func Server() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Define APIs
	e.GET("/v0/docker/version", handler.OutputSistersVersion)

	// Listen at 8080 port
	e.Logger.Fatal(e.Start(":8080"))
}
