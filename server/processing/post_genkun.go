package processing

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo"
	"github.com/ughvj/takamori/drivers"
	"github.com/ughvj/takamori/dml"
	"github.com/ughvj/takamori/types"
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
	defer db.Close()

	if checkAlreadyRegisteredNameKanji(inputData.NameKanji, db) {
		return c.JSON(http.StatusOK, types.NewMessageResponse(inputData.NameKanji + " is already exists."))
	}

	if checkAlreadyRegisteredNameHiragana(inputData.NameYomiHiragana, db) {
		return c.JSON(http.StatusOK, types.NewMessageResponse(inputData.NameYomiHiragana + " is already exists."))
	}


	loadedDML, err := dml.Load("insert_genkun")
	if err != nil {
		return err
	}

	_, err = db.Use().Exec(loadedDML.GetSQL(), inputData.NameKanji, inputData.NameYomiHiragana, inputData.Src)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, types.NewMessageResponse("registered."))
}

func checkAlreadyRegisteredNameKanji(target string, db *drivers.MysqlDriver) bool {
	loadedDML, err := dml.Load("get_genkun_by_name_kanji")
	if err != nil {
		return true
	}

	rows, err := db.Use().Query(loadedDML.GetSQL(), target)
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
	loadedDML, err := dml.Load("get_genkun_by_name_yomi_hiragana")
	if err != nil {
		return true
	}

	rows, err := db.Use().Query(loadedDML.GetSQL(), target)
	if err != nil {
		return true
	}
	defer rows.Close()

	for rows.Next() {
		return true
	}

	return false
}
