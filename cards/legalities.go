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

var LegalitiesNameMap = map[string]string{
	"Standard":         "standard",
	"Future":           "future",
	"Historic":         "historic",
	"Gladiator":        "gladiator",
	"Pioneer":          "pioneer",
	"Explorer":         "explorer",
	"Modern":           "modern",
	"Legacy":           "legacy",
	"Pauper":           "pauper",
	"Vintage":          "vintage",
	"Penny":            "penny",
	"Commander":        "commander",
	"Brawl":            "brawl",
	"Historic Brawl":   "historicbrawl",
	"Alchemy":          "alchemy",
	"Pauper Commander": "paupercommander",
	"Duel":             "duel",
	"Oldschool":        "oldschool",
	"Premodern":        "premodern",
}

var LegalitiesKeysMap = map[string]string{
	"standard":        "Standard",
	"future":          "Future",
	"historic":        "Historic",
	"gladiator":       "Gladiator",
	"pioneer":         "Pioneer",
	"explorer":        "Explorer",
	"modern":          "Modern",
	"legacy":          "Legacy",
	"pauper":          "Pauper",
	"vintage":         "Vintage",
	"penny":           "Penny",
	"commander":       "Commander",
	"brawl":           "Brawl",
	"historicbrawl":   "Historic Brawl",
	"alchemy":         "Alchemy",
	"paupercommander": "Pauper Commander",
	"duel":            "Duel",
	"oldschool":       "Oldschool",
	"premodern":       "Premodern",
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
