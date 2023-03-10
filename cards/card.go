package cards

import (
	"regexp"
	"strings"
)

type Card struct {
	Object          string              `json:"object"`
	Id              string              `json:"id"`
	OracleId        string              `json:"oracle_id"`
	MultiverseIds   []int               `json:"multiverse_ids"`
	MtgoId          int                 `json:"mtgo_id"`
	TcgplayerId     int                 `json:"tcgplayer_id"`
	CardmarketId    int                 `json:"cardmarket_id"`
	Name            string              `json:"name"`
	Lang            string              `json:"lang"`
	ReleasedAt      string              `json:"released_at"`
	Uri             string              `json:"uri"`
	ScryfallUri     string              `json:"scryfall_uri"`
	Layout          string              `json:"layout"`
	HighresImage    bool                `json:"highres_image"`
	ImageStatus     string              `json:"image_status"`
	ImageUris       ImageUris           `json:"image_uris"`
	ManaCost        string              `json:"mana_cost"`
	Cmc             float64             `json:"cmc"`
	TypeLine        string              `json:"type_line"`
	FlavorText      string              `json:"flavor_text"`
	OracleText      string              `json:"oracle_text"`
	Power           string              `json:"power"`
	Toughness       string              `json:"toughness"`
	Colors          []string            `json:"colors"`
	ColorIndicator  []string            `json:"color_indicator"`
	ColorIdentity   []string            `json:"color_identity"`
	Keywords        []interface{}       `json:"keywords"`
	ProducedMana    []string            `json:"produced_mana"`
	CardFaces       []CardFace          `json:"card_faces"`
	AllParts        []AllParts          `json:"all_parts"`
	Legalities      map[string]Legality `json:"legalities"`
	Games           []string            `json:"games"`
	Reserved        bool                `json:"reserved"`
	Foil            bool                `json:"foil"`
	Nonfoil         bool                `json:"nonfoil"`
	Finishes        []string            `json:"finishes"`
	Oversized       bool                `json:"oversized"`
	Promo           bool                `json:"promo"`
	Reprint         bool                `json:"reprint"`
	Variation       bool                `json:"variation"`
	SetId           string              `json:"set_id"`
	Set             string              `json:"set"`
	SetName         string              `json:"set_name"`
	SetType         string              `json:"set_type"`
	SetUri          string              `json:"set_uri"`
	SetSearchUri    string              `json:"set_search_uri"`
	ScryfallSetUri  string              `json:"scryfall_set_uri"`
	RulingsUri      string              `json:"rulings_uri"`
	PrintsSearchUri string              `json:"prints_search_uri"`
	CollectorNumber string              `json:"collector_number"`
	Digital         bool                `json:"digital"`
	Rarity          string              `json:"rarity"`
	CardBackId      string              `json:"card_back_id"`
	Artist          string              `json:"artist"`
	ArtistIds       []string            `json:"artist_ids"`
	IllustrationId  string              `json:"illustration_id"`
	BorderColor     string              `json:"border_color"`
	Frame           string              `json:"frame"`
	FrameEffects    []string            `json:"frame_effects"`
	SecurityStamp   string              `json:"security_stamp"`
	FullArt         bool                `json:"full_art"`
	Textless        bool                `json:"textless"`
	Booster         bool                `json:"booster"`
	StorySpotlight  bool                `json:"story_spotlight"`
	EdhrecRank      int                 `json:"edhrec_rank"`
	Preview         Preview             `json:"preview"`
	Prices          Prices              `json:"prices"`
	RelatedUris     RelatedUris         `json:"related_uris"`
	PurchaseUris    PurchaseUris        `json:"purchase_uris"`
}

var match = regexp.MustCompile(`{.\/?.?}`)

func (c Card) ParseManaCost() [][]string {
	parts := strings.Split(c.ManaCost, "//")
	var res [][]string
	for _, part := range parts {
		set := match.FindAllString(part, -1)
		if len(set) > 0 {
			res = append(res, set)
		}
	}

	return res
}

func (c Card) IsBasicLand() bool {
	return strings.HasPrefix(c.TypeLine, "Basic Land")
}

func (c Card) IsLand() bool {
	return strings.HasPrefix(c.TypeLine, "Land")
}

func (c Card) IsEnchantment() bool {
	return strings.HasPrefix(c.TypeLine, "Enchantment")
}

func (c Card) IsCreature() bool {
	return strings.HasPrefix(c.TypeLine, "Creature")
}

func (c Card) IsLegendaryCreature() bool {
	return strings.HasPrefix(c.TypeLine, "Legendary Creature")
}

func (c Card) IsArtifact() bool {
	return strings.HasPrefix(c.TypeLine, "Artifact")
}

func (c Card) IsLegendaryPlaneswalker() bool {
	return strings.HasPrefix(c.TypeLine, "Legendary Planeswalker")
}

func (c Card) IsInstant() bool {
	return strings.HasPrefix(c.TypeLine, "Instant")
}

func (c Card) IsSorcery() bool {
	return strings.HasPrefix(c.TypeLine, "Sorcery")
}
