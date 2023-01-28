package rulings

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/high-creek-software/goscryfall/endpoint"
	"net/http"
)

type RulingsRepo interface {
	List(url string) ([]Ruling, error)
}

type RestRulingsRepo struct {
	endpoint *endpoint.Endpoint
}

func NewRestRulingsRepo(endpoint *endpoint.Endpoint) RulingsRepo {
	return &RestRulingsRepo{endpoint: endpoint}
}

func (r *RestRulingsRepo) List(url string) ([]Ruling, error) {
	req, err := r.endpoint.NewGetRequestFull(url)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New("error loading rulings")
	}
	var rs endpoint.Response[[]Ruling]
	err = json.NewDecoder(resp.Body).Decode(&rs)
	if err != nil {
		return nil, fmt.Errorf("error parsing rulings: %w", err)
	}

	return rs.Data, nil
}
