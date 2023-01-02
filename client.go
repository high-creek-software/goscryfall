package goscryfall

import (
	"gitlab.com/high-creek-software/goscryfall/bulk"
	"gitlab.com/high-creek-software/goscryfall/cards"
	"gitlab.com/high-creek-software/goscryfall/endpoint"
	"gitlab.com/high-creek-software/goscryfall/sets"
	"gitlab.com/high-creek-software/goscryfall/symbols"
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
}

func NewClient() *Client {
	c := &Client{}
	c.endpoint = endpoint.NewEndpoint(scryfallBaseURL)
	c.SetRepo = sets.NewRestSetRepo(c.endpoint)
	c.CardRepo = cards.NewRestCardRepo(c.endpoint)
	c.SymbolRepo = symbols.NewRestSymbolRepo(c.endpoint)
	c.BulkRepo = bulk.NewRestBulkRepo(c.endpoint)

	return c
}
