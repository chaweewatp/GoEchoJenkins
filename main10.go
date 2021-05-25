package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type AboutData struct {
	ID   string
	Name string
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to my home bro!!!")
	})

	e.GET("/about", aboutHandler)

	e.Logger.Fatal(e.Start(":1323"))
}

func aboutHandler(c echo.Context) error {
	about := AboutData{
		ID:   "400050",
		Name: "John",
	}
	return c.JSON(http.StatusOK, about)
}
