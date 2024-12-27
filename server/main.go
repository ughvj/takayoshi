package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/ughvj/takamori/processing"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo!")
	})

	e.GET("/questions", processing.GetAllQuestions)
	e.POST("/genkun", processing.PostGenkun)

	e.Start(":2434")
}
