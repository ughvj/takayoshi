package types

import (
	"database/sql"
)

type RawQuestions []RawQuestionData

type RawQuestionData struct {
	ID int
	Statement string
	Category string
	CorrectChoice sql.NullBool
	CorrectOrder sql.NullInt32
	GenkunName string
	GenkunSrc string
}

func (q *RawQuestionData) Refs() []interface{} {
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

func (qs *RawQuestions) Processing() []interface{} {
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

func NewTestAllQuestionData() []interface{} {
	rqs := RawQuestions{
		RawQuestionData{
			ID: 1,
			Statement: "初代総理大臣は誰？",
			Category: "choice",
			CorrectChoice: sql.NullBool{Bool: true},
			CorrectOrder: sql.NullInt32{Valid: false},
			GenkunName: "伊藤博文",
			GenkunSrc: "itou_hirobumi.jpg",
		},
		RawQuestionData{
			ID: 1,
			Statement: "初代総理大臣は誰？",
			Category: "choice",
			CorrectChoice: sql.NullBool{Bool: false},
			CorrectOrder: sql.NullInt32{Valid: false},
			GenkunName: "大久保利通",
			GenkunSrc: "okubo_toshimichi.jpg",
		},
		RawQuestionData{
			ID: 1,
			Statement: "初代総理大臣は誰？",
			Category: "choice",
			CorrectChoice: sql.NullBool{Bool: false},
			CorrectOrder: sql.NullInt32{Valid: false},
			GenkunName: "西郷隆盛",
			GenkunSrc: "saigo_takamori.jpg",
		},
		RawQuestionData{
			ID: 1,
			Statement: "初代総理大臣は誰？",
			Category: "choice",
			CorrectChoice: sql.NullBool{Bool: false},
			CorrectOrder: sql.NullInt32{Valid: false},
			GenkunName: "木戸孝允",
			GenkunSrc: "kido_takayoshi.jpg",
		},
		//
		RawQuestionData{
			ID: 2,
			Statement: "一番最初に総理大臣へ就任した年が早い順に選択せよ",
			Category: "order",
			CorrectChoice: sql.NullBool{Valid: false},
			CorrectOrder: sql.NullInt32{Int32: 1},
			GenkunName: "伊藤博文",
			GenkunSrc: "itou_hirobumi.jpg",
		},
		RawQuestionData{
			ID: 2,
			Statement: "一番最初に総理大臣へ就任した年が早い順に選択せよ",
			Category: "order",
			CorrectChoice: sql.NullBool{Valid: false},
			CorrectOrder: sql.NullInt32{Int32: 2},
			GenkunName: "黒田清隆",
			GenkunSrc: "kuroda_kiyotaka.jpg",
		},
		RawQuestionData{
			ID: 2,
			Statement: "一番最初に総理大臣へ就任した年が早い順に選択せよ",
			Category: "order",
			CorrectChoice: sql.NullBool{Valid: false},
			CorrectOrder: sql.NullInt32{Int32: 3},
			GenkunName: "山縣有朋",
			GenkunSrc: "yamagata_aritomo.jpg",
		},
		RawQuestionData{
			ID: 2,
			Statement: "一番最初に総理大臣へ就任した年が早い順に選択せよ",
			Category: "order",
			CorrectChoice: sql.NullBool{Valid: false},
			CorrectOrder: sql.NullInt32{Int32: 4},
			GenkunName: "松方正義",
			GenkunSrc: "matsukata_masayoshi.jpg",
		},
	}
	return rqs.Processing()
}
