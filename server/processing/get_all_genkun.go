package processing

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/ughvj/takayoshi/drivers"
	"github.com/ughvj/takayoshi/dml"
	"github.com/ughvj/takayoshi/types"
)

func GetAllGenkunDryrun(c echo.Context) error {
	return c.JSON(http.StatusOK, types.NewTestAllGenkunData())
}

func GetAllGenkun(c echo.Context) error {
	db, err := drivers.NewMysqlDriver()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	loadedDML, err := dml.Loader.Get("get_all_genkun")

	rows, err := db.Use().Query(loadedDML)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var gs types.Genkuns
	for rows.Next() {
		var g types.Genkun
		err := rows.Scan(g.Refs()...)
		if err != nil {
			panic(err.Error())
		}
		gs = append(gs, g)
	}

	return c.JSON(http.StatusOK, gs)
}
