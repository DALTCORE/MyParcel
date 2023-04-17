package myparcel

import (
	"net/http"
	"sync"
)

// Config contains the configuration for the MyParcel client.
type Config struct {
	APIKey string
}

// Client is the MyParcel client.
type Client struct {
	apiBaseURL string
	config     Config
	httpClient *http.Client
	sync.Mutex
}

// NewClient returns a new MyParcel client.
// The API key is required to use the MyParcel API.
func NewClient(c Config) *Client {
	return &Client{
		apiBaseURL: "https://api.myparcel.nl",
		httpClient: &http.Client{},
		config:     c,
	}
}
