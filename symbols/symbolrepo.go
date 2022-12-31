package symbols

import (
	"encoding/json"
	"fmt"
	"gitlab.com/high-creek-software/goscryfall/endpoint"
	"net/http"
)

const (
	symbols = "/symbology"
)

type SymbolRepo interface {
	ListSymbols() (*endpoint.Response[[]Symbol], error)
}

type RestSymbolRepo struct {
	endpoint *endpoint.Endpoint
}

func NewRestSymbolRepo(endpoint *endpoint.Endpoint) SymbolRepo {
	return &RestSymbolRepo{endpoint: endpoint}
}

func (r *RestSymbolRepo) ListSymbols() (*endpoint.Response[[]Symbol], error) {
	req, err := r.endpoint.NewGetRequest(symbols)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("could not load symbols")
	}

	var s endpoint.Response[[]Symbol]
	err = json.NewDecoder(resp.Body).Decode(&s)
	if err != nil {
		return nil, fmt.Errorf("error parsing symbols %w", err)
	}

	return &s, nil
}
