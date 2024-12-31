package processing

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/ughvj/takayoshi/types"
)

func PostQuestionDryrun(c echo.Context) error {
	var inputData types.InputDataPostQuestion
	if err := c.Bind(&inputData); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, types.NewMessageResponse("post /question is dryrunning. statement: " + inputData.Statement))
}

func PostQuestion(c echo.Context) error {
	return c.JSON(http.StatusOK, types.NewMessageResponse("post_question"))
}
