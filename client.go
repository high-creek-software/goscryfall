package goscryfall

import (
	"github.com/high-creek-software/goscryfall/bulk"
	"github.com/high-creek-software/goscryfall/cards"
	"github.com/high-creek-software/goscryfall/endpoint"
	"github.com/high-creek-software/goscryfall/rulings"
	"github.com/high-creek-software/goscryfall/sets"
	"github.com/high-creek-software/goscryfall/symbols"
)

const (
	scryfallBaseURL = "https://api.scryfall.com"
)

type Client struct {
	endpoint *endpoint.Endpoint

	sets.SetRepo
	cards.CardRepo
	symbols.SymbolRepo
	bulk.BulkRepo
	rulings.RulingsRepo
}

func NewClient() *Client {
	c := &Client{}
	c.endpoint = endpoint.NewEndpoint(scryfallBaseURL)
	c.SetRepo = sets.NewRestSetRepo(c.endpoint)
	c.CardRepo = cards.NewRestCardRepo(c.endpoint)
	c.SymbolRepo = symbols.NewRestSymbolRepo(c.endpoint)
	c.BulkRepo = bulk.NewRestBulkRepo(c.endpoint)
	c.RulingsRepo = rulings.NewRestRulingsRepo(c.endpoint)

	return c
}
