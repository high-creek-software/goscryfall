package bulk

import "time"

type Bulk struct {
	Object          string    `json:"object"`
	Id              string    `json:"id"`
	Type            string    `json:"type"`
	UpdatedAt       time.Time `json:"updated_at"`
	Uri             string    `json:"uri"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	Size            int       `json:"size"`
	DownloadUri     string    `json:"download_uri"`
	ContentType     string    `json:"content_type"`
	ContentEncoding string    `json:"content_encoding"`
}
