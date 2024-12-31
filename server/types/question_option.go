package types

import (
	"database/sql"
	"strings"
	"strconv"
)

type QuestionOptions []QuestionOption

func (qos *QuestionOptions) GenerateBulkSentence() string {
	var processed []string
	for _, qo := range *qos {
		var elements []string
		if (qo.CorrectChoice.Valid) {
			if (qo.CorrectChoice.Bool) {
				elements = append(elements, "true")
			} else {
				elements = append(elements, "false")
			}
		} else {
			elements = append(elements, "null")
		}

		if (qo.CorrectOrder.Valid) {
			elements = append(elements, string(qo.CorrectOrder.Int32))
		} else {
			elements = append(elements, "null")
		}

		elements = append(elements, strconv.Itoa(qo.QuestionId))
		elements = append(elements, strconv.Itoa(qo.GenkunId))

		processed = append(processed, "(" + strings.Join(elements, ",") + ")")
	}

	return strings.Join(processed, ",")
}

type QuestionOption struct {
	Id int
	CorrectChoice sql.NullBool
	CorrectOrder sql.NullInt32
	QuestionId int
	GenkunId int
}
