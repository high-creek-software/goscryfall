package cards

type Prices struct {
	Usd       string      `json:"usd"`
	UsdFoil   string      `json:"usd_foil"`
	UsdEtched interface{} `json:"usd_etched"`
	Eur       string      `json:"eur"`
	EurFoil   string      `json:"eur_foil"`
	Tix       string      `json:"tix"`
}
