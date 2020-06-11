package main

import (
	"github.com/labstack/echo/middleware"
	router "github.com/rafaelarodrigues90/usuarios/routers"
)

func main() {
	e := router.App

	e.Use(middleware.Logger())

	e.Logger.Fatal(e.Start(":3000"))
}
