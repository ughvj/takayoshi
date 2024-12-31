package types

import (
	"database/sql"
)

//
type InputDataPostQuestion struct {
	Statement string `json:"statement"`
	Category string `json:"category"`
	Options []InputDataPostQuestionOptions `json:"options"`
}

type InputDataPostQuestionOptions struct {
	GenkunID int `json:"genkun_id"`
	Correct interface{} `json:"correct"`
}

//
type QueryGetAllQuestion []QueryGetAllQuestionOne

type QueryGetAllQuestionOne struct {
	ID int
	Statement string
	Category string
	CorrectChoice sql.NullBool
	CorrectOrder sql.NullInt32
	GenkunName string
	GenkunSrc string
}

func (q *QueryGetAllQuestionOne) Refs() []interface{} {
	return []interface{}{
		&q.ID,
		&q.Statement,
		&q.Category,
		&q.CorrectChoice,
		&q.CorrectOrder,
		&q.GenkunName,
		&q.GenkunSrc,
	}
}

func (qs *QueryGetAllQuestion) GenerateResponseData() []interface{} {
	var processed []interface{}

	type ChoiceQuestionOption struct {
		Correct bool `json:"correct"`
		GenkunName string `json:"caption"`
		GenkunSrc string `json:"src"`
	}
	var cqos []ChoiceQuestionOption

	type OrderQuestionOption struct {
		Correct int32 `json:"correct"`
		GenkunName string `json:"caption"`
		GenkunSrc string `json:"src"`
	}
	var oqos []OrderQuestionOption

	for i, q := range *qs {
		switch q.Category {
		case "choice":
			cqos = append(cqos, ChoiceQuestionOption{
				q.CorrectChoice.Bool,
				q.GenkunName,
				q.GenkunSrc,
			})
		case "order":
			oqos = append(oqos, OrderQuestionOption{
				q.CorrectOrder.Int32,
				q.GenkunName,
				q.GenkunSrc,
			})
		}
		if (i+1) % 4 == 0 {
			switch q.Category {
			case "choice":
				type ChoiceQuestion struct {
					ID int `json:"id"`
					Statement string `json:"statement"`
					Category string `json:"category"`
					Options []ChoiceQuestionOption `json:"options"`
				}
				processed = append(processed, ChoiceQuestion{
					q.ID,
					q.Statement,
					q.Category,
					cqos,
				})
				cqos = []ChoiceQuestionOption{}
			case "order":
				type OrderQuestion struct {
					ID int `json:"id"`
					Statement string `json:"statement"`
					Category string `json:"category"`
					Options []OrderQuestionOption `json:"options"`
				}
				processed = append(processed, OrderQuestion{
					q.ID,
					q.Statement,
					q.Category,
					oqos,
				})
				oqos = []OrderQuestionOption{}
			}
		}
	}

	return processed
}
