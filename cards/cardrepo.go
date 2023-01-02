package cards

import (
	"encoding/json"
	"errors"
	"fmt"
	"gitlab.com/high-creek-software/goscryfall/endpoint"
	"net/http"
	"net/url"
)

const (
	cardSearchRoute = "/cards/search"
)

type CardRepo interface {
	ListCards(setID string, next string) (*endpoint.Response[[]Card], error)
}

type RestCardRepo struct {
	endpoint *endpoint.Endpoint
}

func (r RestCardRepo) ListCards(setID string, next string) (*endpoint.Response[[]Card], error) {
	var req *http.Request
	var err error

	if next == "" {
		req, err = r.endpoint.NewGetRequest(cardSearchRoute)
		if err != nil {
			return nil, err
		}

		q := url.Values{}
		q.Add("q", fmt.Sprintf("set:%s", setID))

		req.URL.RawQuery = q.Encode()
	} else {
		req, err = r.endpoint.NewGetRequestFull(next)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("error loading cards for set: %s", setID))
	}

	var cards endpoint.Response[[]Card]
	err = json.NewDecoder(resp.Body).Decode(&cards)
	if err != nil {
		return nil, fmt.Errorf("error parsing cards: %w", err)
	}

	return &cards, nil
}

func NewRestCardRepo(endpoint *endpoint.Endpoint) CardRepo {
	return &RestCardRepo{endpoint: endpoint}
}
