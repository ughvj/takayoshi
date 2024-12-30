package processing

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init(isDryrun bool) *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{ "http://localhost:3000" },
		AllowHeaders: []string{ echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept },
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo!")
	})

	if isDryrun {
		e.GET("/questions", GetAllQuestionsDryrun)
		e.POST("/genkun", PostGenkunDryrun)
	} else {
		e.GET("/questions", GetAllQuestions)
		e.POST("/genkun", PostGenkun)
	}

	return e;
}
