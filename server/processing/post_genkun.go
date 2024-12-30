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
	inputGenkunNameKanji := c.FormValue("genkun_name_kanji")
	inputGenkunNameHiragana := c.FormValue("genkun_name_hiragana")
	inputGenkunSrc := c.FormValue("genkun_src")

	template := "genkun is dryrunning. genkun_name_kanji: %s, genkun_name_hiragana: %s genkun_src: %s"

	return c.JSON(
		http.StatusOK,
		types.NewMessageResponse(
			fmt.Sprintf(
				template,
				inputGenkunNameKanji,
				inputGenkunNameHiragana,
				inputGenkunSrc,
			),
		),
	)
}

func PostGenkun(c echo.Context) error {
	inputGenkunNameKanji := c.FormValue("genkun_name_kanji")
	inputGenkunNameHiragana := c.FormValue("genkun_name_hiragana")
	inputGenkunSrc := c.FormValue("genkun_src")

	db, err := drivers.NewMysqlDriver()
	if err != nil {
		return err
	}
	defer db.Close()

	if checkAlreadyRegisteredNameKanji(inputGenkunNameKanji, db) {
		return c.JSON(http.StatusOK, types.NewMessageResponse(inputGenkunNameKanji + " is already exists."))
	}

	if checkAlreadyRegisteredNameHiragana(inputGenkunNameHiragana, db) {
		return c.JSON(http.StatusOK, types.NewMessageResponse(inputGenkunNameHiragana + " is already exists."))
	}


	loadedDML, err := dml.Load("insert_genkun")
	if err != nil {
		return err
	}

	_, err = db.Use().Exec(loadedDML.GetSQL(), inputGenkunNameKanji, inputGenkunNameHiragana, inputGenkunSrc)
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
