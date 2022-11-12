package main

import (
	"net/http"
	passGenerator "password-generator-api/functions"
	"strconv"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Ok!")
	})
	e.GET("/api", processor)
	e.Logger.Fatal(e.Start(":3000"))
}

type Body struct {
	Message string `json:"message" xml:"message"`
}

func processor(c echo.Context) error {
	lenght, _ := strconv.Atoi(c.QueryParam("lenght"))
	difficulty, _ := strconv.Atoi(c.QueryParam("difficulty"))
	newPass := passGenerator.PassGen(lenght, difficulty)
	p := &Body{
		Message: newPass,
	}
	return c.JSON(http.StatusOK, p)
}
