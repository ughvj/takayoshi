package processing

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/ughvj/takayoshi/drivers"
	"github.com/ughvj/takayoshi/dml"
	"github.com/ughvj/takayoshi/types"
)

func GetAllQuestionsDryrun(c echo.Context) error {
	return c.JSON(http.StatusOK, types.NewTestAllQuestionData())
}

func GetAllQuestions(c echo.Context) error {
	db, err := drivers.NewMysqlDriver()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	loadedDML, err := dml.Loader.Get("get_all_question")

	rows, err := db.Use().Query(loadedDML)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var qs types.QueryGetAllQuestion
	for rows.Next() {
		var q types.QueryGetAllQuestionOne
		err := rows.Scan(q.Refs()...)
		if err != nil {
			panic(err.Error())
		}
		qs = append(qs, q)
	}

	return c.JSON(http.StatusOK, qs.GenerateResponseData())
}
