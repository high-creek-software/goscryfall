package rulings

type Ruling struct {
	Object      string `json:"object"`
	OracleId    string `json:"oracle_id"`
	Source      string `json:"source"`
	PublishedAt string `json:"published_at"`
	Comment     string `json:"comment"`
}
