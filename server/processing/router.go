package processing

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/ughvj/takayoshi/config"
)

func Init() *echo.Echo {
	e := echo.New()

	conf := config.Loader.Get()
	fmt.Printf("%+v\n", conf)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: conf.AllowOrigins,
		AllowHeaders: []string{ echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept },
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo!")
	})

	if conf.Dryrun {
		fmt.Printf("Running on dryrun.")
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
