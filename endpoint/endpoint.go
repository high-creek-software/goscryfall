package endpoint

import (
	"net/http"
)

const (
	headerAccept      = "Accept"
	headerContentType = "Content-Type"
	applicationJson   = "application/json"
)

type Endpoint struct {
	baseURL string
}

func NewEndpoint(baseURL string) *Endpoint {
	return &Endpoint{baseURL: baseURL}
}

func (e *Endpoint) NewGetRequest(resource string) (*http.Request, error) {
	url := e.baseURL + resource

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add(headerAccept, applicationJson)

	return req, nil
}

func (e *Endpoint) NewGetRequestFull(uri string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add(headerAccept, applicationJson)

	return req, nil
}