package main

import (
	"github.com/labstack/echo/v4"
	"password-generator-api/routes"
)

func main() {
	e := echo.New()
	e.Static("/", "views")
	e.GET("/health", routes.Health)
	e.GET("/gen", routes.Generate)
	e.Logger.Fatal(e.Start(":3000"))
}
