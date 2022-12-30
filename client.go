package goscryfall

import (
	"gitlab.com/high-creek-software/goscryfall/endpoint"
	"gitlab.com/high-creek-software/goscryfall/sets"
)

const (
	scryfallBaseURL = "https://api.scryfall.com"
)

type Client struct {
	endpoint *endpoint.Endpoint

	sets.SetRepo
}

func NewClient() *Client {
	c := &Client{}
	c.endpoint = endpoint.NewEndpoint(scryfallBaseURL)
	c.SetRepo = sets.NewRestSetRepo(c.endpoint)

	return c
}
