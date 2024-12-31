package processing

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"
	"errors"
	"fmt"

	"github.com/labstack/echo"
	"github.com/ughvj/takayoshi/types"
	"github.com/ughvj/takayoshi/drivers"
	"github.com/ughvj/takayoshi/dml"
)

func PostQuestionDryrun(c echo.Context) error {
	var inputData types.InputDataPostQuestion
	if err := c.Bind(&inputData); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, types.NewMessageResponse("post /question is dryrunning. statement: " + inputData.Statement))
}

func PostQuestion(c echo.Context) error {
	var inputData types.InputDataPostQuestion
	if err := c.Bind(&inputData); err != nil {
		return err
	}

	db, err := drivers.NewMysqlDriver()
	if err != nil {
		return err
	}
	defer db.Use().Close()

	tx, err := db.Use().Begin()
	if err != nil {
		return err
	}

	// 1. まずはquestionsテーブルへ登録
	loadedDML, _ := dml.Loader.Get("insert_question")
	inserted, err := tx.Exec(loadedDML, inputData.Statement, inputData.Category)
	if err != nil {
		tx.Rollback()
		return err
	}
	insertedQuestionId, err := inserted.LastInsertId()

	// 2. 次に、クイズに使われている元勲のデータを全て取得しにいく
	var genkunIds []string
	for _, opt := range inputData.Options {
		genkunIds = append(genkunIds, strconv.Itoa(opt.GenkunID))
	}
	loadedDML, _ = dml.Loader.EmbedAndGet("get_genkun_by_ids_embedded", "(" + strings.Join(genkunIds, ",") + ")")
	fmt.Printf("%s", loadedDML)
	rows, err := db.Use().Query(loadedDML)
	if err != nil {
		tx.Rollback()
		return err
	}

	var gs types.Genkuns
	for rows.Next() {
		var g types.Genkun
		err := rows.Scan(g.Refs()...)
		if err != nil {
			tx.Rollback()
			return err
		}
		gs = append(gs, g)
	}

	// 3. 2.をquestion_option構造体へマッピングする
	var qos types.QuestionOptions
	for _, opt := range inputData.Options {
		var qo types.QuestionOption
		qo.QuestionId = int(insertedQuestionId)
		qo.GenkunId = opt.GenkunID
		switch inputData.Category {
		case "choice":
			v, ok := opt.Correct.(bool)
			if ok {
				qo.CorrectChoice = sql.NullBool{Valid: true, Bool: v}
				qo.CorrectOrder = sql.NullInt32{Valid: false}
			} else {
				tx.Rollback()
				return errors.New("failed: bool converted")
			}
		case "order":
			v, ok := opt.Correct.(int32)
			if ok {
				qo.CorrectChoice = sql.NullBool{Valid: false}
				qo.CorrectOrder = sql.NullInt32{Valid: true, Int32: v}
			} else {
				tx.Rollback()
				return errors.New("failed: int32 converted")
			}
		}

		qos = append(qos, qo)
	}

	// 4. 最後に、3.をもとにquestion_optionsテーブルへ登録
	loadedDML, _ = dml.Loader.EmbedAndGet("bulk_insert_question_option_embedded", qos.GenerateBulkSentence())
	_, err = tx.Exec(loadedDML)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	return c.JSON(http.StatusOK, types.NewMessageResponse("registered."))
}
