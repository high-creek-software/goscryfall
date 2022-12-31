package symbols

type Symbol struct {
	Object             string      `json:"object"`
	Symbol             string      `json:"symbol"`
	SvgUri             string      `json:"svg_uri"`
	LooseVariant       interface{} `json:"loose_variant"`
	English            string      `json:"english"`
	Transposable       bool        `json:"transposable"`
	RepresentsMana     bool        `json:"represents_mana"`
	AppearsInManaCosts bool        `json:"appears_in_mana_costs"`
	Cmc                float64     `json:"cmc"`
	Funny              bool        `json:"funny"`
	Colors             []string    `json:"colors"`
	GathererAlternates interface{} `json:"gatherer_alternates"`
}
