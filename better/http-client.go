package main

import (
	"io/ioutil"
	"net/http"
)

// IHTTPClient interface to wrap http functionality
type IHTTPClient interface {
	Get(url string) (*http.Response, error)
}

// IWebClient interface to wrap all functionality in http-client
type IWebClient interface {
	GetContent(string) (string, error)
}

// WebClient contains the http client so we can mock this
type WebClient struct {
	c IHTTPClient
}

// NewWebClient created an instance of IWebClient
func NewWebClient() IWebClient {
	return WebClient{
		c: &http.Client{},
	}
}

// GetContent retreive the content from an URL
func (w WebClient) GetContent(u string) (string, error) {
	resp, err := w.c.Get(u)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

//go:generate mockgen -source=http-client.go -package=main -destination=./mocks/http-client_mock.go -package=mocks
