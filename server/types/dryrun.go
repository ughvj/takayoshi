package types

import (
	"database/sql"
)

func NewTestAllGenkunData() Genkuns {
	return Genkuns{
		{ 1, "itou_hirobumi.jpg", "伊藤博文", "いとうひろぶみ" },
		{ 2, "ookubo_toshimichi.jpg", "大久保利通", "おおくぼとしみち" },
		{ 3, "saigou_takamori.jpg", "西郷隆盛", "さいごうたかもり" },
		{ 4, "kido_takayoshi", "木戸孝允", "きどたかよし" },
	}
}

func NewTestAllQuestionData() []interface{} {
	rqs := QueryGetAllQuestion{
		QueryGetAllQuestionOne{
			ID: 1,
			Statement: "初代総理大臣は誰？",
			Category: "choice",
			CorrectChoice: sql.NullBool{Bool: true},
			CorrectOrder: sql.NullInt32{Valid: false},
			GenkunName: "伊藤博文",
			GenkunSrc: "itou_hirobumi.jpg",
		},
		QueryGetAllQuestionOne{
			ID: 1,
			Statement: "初代総理大臣は誰？",
			Category: "choice",
			CorrectChoice: sql.NullBool{Bool: false},
			CorrectOrder: sql.NullInt32{Valid: false},
			GenkunName: "大久保利通",
			GenkunSrc: "ookubo_toshimichi.jpg",
		},
		QueryGetAllQuestionOne{
			ID: 1,
			Statement: "初代総理大臣は誰？",
			Category: "choice",
			CorrectChoice: sql.NullBool{Bool: false},
			CorrectOrder: sql.NullInt32{Valid: false},
			GenkunName: "西郷隆盛",
			GenkunSrc: "saigou_takamori.jpg",
		},
		QueryGetAllQuestionOne{
			ID: 1,
			Statement: "初代総理大臣は誰？",
			Category: "choice",
			CorrectChoice: sql.NullBool{Bool: false},
			CorrectOrder: sql.NullInt32{Valid: false},
			GenkunName: "木戸孝允",
			GenkunSrc: "kido_takayoshi.jpg",
		},
		//
		QueryGetAllQuestionOne{
			ID: 2,
			Statement: "一番最初に総理大臣へ就任した年が早い順に選択せよ",
			Category: "order",
			CorrectChoice: sql.NullBool{Valid: false},
			CorrectOrder: sql.NullInt32{Int32: 1},
			GenkunName: "伊藤博文",
			GenkunSrc: "itou_hirobumi.jpg",
		},
		QueryGetAllQuestionOne{
			ID: 2,
			Statement: "一番最初に総理大臣へ就任した年が早い順に選択せよ",
			Category: "order",
			CorrectChoice: sql.NullBool{Valid: false},
			CorrectOrder: sql.NullInt32{Int32: 2},
			GenkunName: "黒田清隆",
			GenkunSrc: "kuroda_kiyotaka.jpg",
		},
		QueryGetAllQuestionOne{
			ID: 2,
			Statement: "一番最初に総理大臣へ就任した年が早い順に選択せよ",
			Category: "order",
			CorrectChoice: sql.NullBool{Valid: false},
			CorrectOrder: sql.NullInt32{Int32: 3},
			GenkunName: "山縣有朋",
			GenkunSrc: "yamagata_aritomo.jpg",
		},
		QueryGetAllQuestionOne{
			ID: 2,
			Statement: "一番最初に総理大臣へ就任した年が早い順に選択せよ",
			Category: "order",
			CorrectChoice: sql.NullBool{Valid: false},
			CorrectOrder: sql.NullInt32{Int32: 4},
			GenkunName: "松方正義",
			GenkunSrc: "matsukata_masayoshi.jpg",
		},
	}
	return rqs.GenerateResponseData()
}
