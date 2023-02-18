package routes

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"password-generator-api/utils"
	"strconv"
)

type Error struct {
	Error string `json:"error"`
}

type Message struct {
	Message string `json:"password"`
}

func Generate(c echo.Context) error {
	if c.QueryParam("length") == "" {
		res := &Error{
			Error: "Provides the length of the password to be returned.",
		}
		err := c.JSON(http.StatusOK, res)
		if err != nil {
			return err
		}
	} else {
		length, _ := strconv.Atoi(c.QueryParam("length"))
		pass := utils.Gen(length)

		res := &Message{
			Message: pass,
		}
		err := c.JSON(http.StatusOK, res)
		if err != nil {
			return err
		}
	}
	return nil
}
