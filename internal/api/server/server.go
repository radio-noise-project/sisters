package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Start() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	router(e)
	e.Logger.Fatal(e.Start(":10032"))
}
