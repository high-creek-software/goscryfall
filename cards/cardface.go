package cards

type CardFace struct {
	Object         string    `json:"object"`
	Name           string    `json:"name"`
	FlavorName     string    `json:"flavor_name"`
	ManaCost       string    `json:"mana_cost"`
	TypeLine       string    `json:"type_line"`
	OracleText     string    `json:"oracle_text"`
	Colors         []string  `json:"colors"`
	ColorIndicator []string  `json:"color_indicator"`
	Loyalty        string    `json:"loyalty"`
	Artist         string    `json:"artist"`
	ArtistId       string    `json:"artist_id"`
	IllustrationId string    `json:"illustration_id"`
	ImageUris      ImageUris `json:"image_uris"`
	Power          string    `json:"power"`
	Toughness      string    `json:"toughness"`
}
