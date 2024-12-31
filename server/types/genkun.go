package types

type Genkuns []Genkun

func (gs *Genkuns) findById(id int) *Genkun {
	for _, g := range *gs {
		if g.Id == id {
			return &g
		}
	}
	return nil
}

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
