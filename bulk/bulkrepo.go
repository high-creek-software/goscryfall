package bulk

import (
	"encoding/json"
	"errors"
	"fmt"
	"gitlab.com/high-creek-software/goscryfall/endpoint"
	"net/http"
)

const (
	bulkData = "/bulk-data"
)

type BulkRepo interface {
	ListBulk() (*endpoint.Response[[]Bulk], error)
}

type RestBulkRepo struct {
	endpoint *endpoint.Endpoint
}

func NewRestBulkRepo(endpoint *endpoint.Endpoint) BulkRepo {
	return &RestBulkRepo{endpoint: endpoint}
}

func (r *RestBulkRepo) ListBulk() (*endpoint.Response[[]Bulk], error) {
	req, err := r.endpoint.NewGetRequest(bulkData)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("error loading bulk data: %d", resp.StatusCode))
	}

	var bulk endpoint.Response[[]Bulk]
	err = json.NewDecoder(resp.Body).Decode(&bulk)
	if err != nil {
		return nil, fmt.Errorf("error parsing bulk data: %w", err)
	}

	return &bulk, nil
}
