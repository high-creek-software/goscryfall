package endpoint

type Response[T any] struct {
	Object   string `json:"object"`
	HasMore  bool   `json:"has_more"`
	NextPage string `json:"next_page"`
	Data     T      `json:"data"`
}
