package cards

import (
	"encoding/json"
	"strings"
)

//type Legalitites struct {
//	Standard        Legality `json:"standard"`
//	Future          Legality `json:"future"`
//	Historic        Legality `json:"historic"`
//	Gladiator       Legality `json:"gladiator"`
//	Pioneer         Legality `json:"pioneer"`
//	Explorer        Legality `json:"explorer"`
//	Modern          Legality `json:"modern"`
//	Legacy          Legality `json:"legacy"`
//	Pauper          Legality `json:"pauper"`
//	Vintage         Legality `json:"vintage"`
//	Penny           Legality `json:"penny"`
//	Commander       Legality `json:"commander"`
//	Brawl           Legality `json:"brawl"`
//	Historicbrawl   Legality `json:"historicbrawl"`
//	Alchemy         Legality `json:"alchemy"`
//	Paupercommander Legality `json:"paupercommander"`
//	Duel            Legality `json:"duel"`
//	Oldschool       Legality `json:"oldschool"`
//	Premodern       Legality `json:"premodern"`
//}

const (
	Standard        = "standard"
	Future          = "future"
	Historic        = "historic"
	Gladiator       = "gladiator"
	Pioneer         = "pioneer"
	Explorer        = "explorer"
	Modern          = "modern"
	Legacy          = "legacy"
	Pauper          = "pauper"
	Vintage         = "vintage"
	Penny           = "penny"
	Commander       = "commander"
	Brawl           = "brawl"
	HistoricBrawl   = "historicbrawl"
	Alchemy         = "alchemy"
	PauperCommander = "paupercommander"
	Duel            = "duel"
	Oldschool       = "oldschool"
	Premodern       = "premodern"
)

var LegalitiesNameMap = map[string]string{
	"Standard":         Standard,
	"Future":           Future,
	"Historic":         Historic,
	"Gladiator":        Gladiator,
	"Pioneer":          Pioneer,
	"Explorer":         Explorer,
	"Modern":           Modern,
	"Legacy":           Legacy,
	"Pauper":           Pauper,
	"Vintage":          Vintage,
	"Penny":            Penny,
	"Commander":        Commander,
	"Brawl":            Brawl,
	"Historic Brawl":   HistoricBrawl,
	"Alchemy":          Alchemy,
	"Pauper Commander": PauperCommander,
	"Duel":             Duel,
	"Oldschool":        Oldschool,
	"Premodern":        Premodern,
}

var LegalitiesKeysMap = map[string]string{
	Standard:        "Standard",
	Future:          "Future",
	Historic:        "Historic",
	Gladiator:       "Gladiator",
	Pioneer:         "Pioneer",
	Explorer:        "Explorer",
	Modern:          "Modern",
	Legacy:          "Legacy",
	Pauper:          "Pauper",
	Vintage:         "Vintage",
	Penny:           "Penny",
	Commander:       "Commander",
	Brawl:           "Brawl",
	HistoricBrawl:   "Historic Brawl",
	Alchemy:         "Alchemy",
	PauperCommander: "Pauper Commander",
	Duel:            "Duel",
	Oldschool:       "Oldschool",
	Premodern:       "Premodern",
}

type Legality int

func (l *Legality) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.String())
}

func (l *Legality) UnmarshalJSON(bs []byte) error {
	str := strings.Trim(string(bs), `"`)
	*l = LegalityFromString(str)
	return nil
}

var serializedLegality = []string{"legal", "not_legal", "restricted", "banned", "unknown"}
var displayLegality = []string{"Legal", "Not Legal", "Restricted", "Banned", "Unknown"}

const (
	Legal Legality = iota
	NotLegal
	Restricted
	Banned
	Unknown
)

func (l Legality) Display() string {
	return displayLegality[l]
}

func (l Legality) String() string {
	return serializedLegality[l]
}

func LegalityFromString(l string) Legality {
	for i, s := range serializedLegality {
		if s == l {
			return Legality(i)
		}
	}
	return Unknown
}
