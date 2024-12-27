package processing

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/ughvj/takamori/drivers"
	"github.com/ughvj/takamori/dml"
	"github.com/ughvj/takamori/types"
)

func GetAllQuestions(c echo.Context) error {
	db, err := drivers.NewMysqlDriver()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	loadedDML, err := dml.Load("get_all_question")

	rows, err := db.Use().Query(loadedDML.GetSQL())
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var qs types.RawQuestions
	for rows.Next() {
		var q types.RawQuestionData
		err := rows.Scan(q.Refs()...)
		if err != nil {
			panic(err.Error())
		}
		qs = append(qs, q)
	}

	return c.JSON(http.StatusOK, qs.Processing())
}
