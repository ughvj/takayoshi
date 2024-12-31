package types

type Genkuns []Genkun

type Genkun struct {
	Id int `json:"id"`
	Src string `json:"src"`
	NameKanji string `json:"name_kanji"`
	NameYomiHiragana string `json:"name_yomi_hiragana"`
}

func (g *Genkun) Refs() []interface{} {
	return []interface{}{
		&g.Id,
		&g.Src,
		&g.NameKanji,
		&g.NameYomiHiragana,
	}
}

type InputDataPostGenkun struct {
	Src string `json:"src"`
	NameKanji string `json:"name_kanji"`
	NameYomiHiragana string `json:"name_yomi_hiragana"`
}

func NewTestAllGenkunData() Genkuns {
	return Genkuns{
		{ 1, "itou_hirobumi.jpg", "伊藤博文", "いとうひろぶみ" },
		{ 2, "ookubo_toshimichi.jpg", "大久保利通", "おおくぼとしみち" },
		{ 3, "saigou_takamori.jpg", "西郷隆盛", "さいごうたかもり" },
		{ 4, "kido_takayoshi", "木戸孝允", "きどたかよし" },
	}
}
