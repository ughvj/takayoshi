package processing

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo"
	"github.com/ughvj/takayoshi/drivers"
	"github.com/ughvj/takayoshi/dml"
	"github.com/ughvj/takayoshi/types"
)

func PostGenkunDryrun(c echo.Context) error {
	var inputData types.InputDataPostGenkun
	if err := c.Bind(&inputData); err != nil {
		return err
	}

	template := "post /genkun is dryrunning. name_kanji: %s, name_yomi_hiragana: %s src: %s"

	return c.JSON(
		http.StatusOK,
		types.NewMessageResponse(
			fmt.Sprintf(
				template,
				inputData.NameKanji,
				inputData.NameYomiHiragana,
				inputData.Src,
			),
		),
	)
}

func PostGenkun(c echo.Context) error {
	var inputData types.InputDataPostGenkun
	if err := c.Bind(&inputData); err != nil {
		return err
	}

	db, err := drivers.NewMysqlDriver()
	if err != nil {
		return err
	}
	defer db.Use().Close()

	if checkAlreadyRegisteredNameKanji(inputData.NameKanji, db) {
		return c.JSON(http.StatusOK, types.NewMessageResponse(inputData.NameKanji + " is already exists."))
	}

	if checkAlreadyRegisteredNameHiragana(inputData.NameYomiHiragana, db) {
		return c.JSON(http.StatusOK, types.NewMessageResponse(inputData.NameYomiHiragana + " is already exists."))
	}


	loadedDML, err := dml.Loader.Get("insert_genkun")
	if err != nil {
		return err
	}

	_, err = db.Use().Exec(loadedDML, inputData.NameKanji, inputData.NameYomiHiragana, inputData.Src)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, types.NewMessageResponse("registered."))
}

func checkAlreadyRegisteredNameKanji(target string, db *drivers.MysqlDriver) bool {
	loadedDML, err := dml.Loader.Get("get_genkun_by_name_kanji")
	if err != nil {
		return true
	}

	rows, err := db.Use().Query(loadedDML, target)
	if err != nil {
		return true
	}
	defer rows.Close()

	for rows.Next() {
		return true
	}

	return false
}

func checkAlreadyRegisteredNameHiragana(target string, db *drivers.MysqlDriver) bool {
	loadedDML, err := dml.Loader.Get("get_genkun_by_name_yomi_hiragana")
	if err != nil {
		return true
	}

	rows, err := db.Use().Query(loadedDML, target)
	if err != nil {
		return true
	}
	defer rows.Close()

	for rows.Next() {
		return true
	}

	return false
}
