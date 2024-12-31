package processing

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/ughvj/takayoshi/config"
)

func Init(isDryrun bool) *echo.Echo {
	e := echo.New()

	gl := config.Loader.Get()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: gl.AllowOrigins,
		AllowHeaders: []string{ echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept },
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo!")
	})

	if isDryrun {
		e.GET("/questions", GetAllQuestionsDryrun)
		e.POST("/genkun", PostGenkunDryrun)
		e.POST("/question", PostQuestionDryrun)
	} else {
		e.GET("/questions", GetAllQuestions)
		e.POST("/genkun", PostGenkun)
		e.POST("/question", PostQuestion)
	}

	return e;
}
