package reservoir

import (
	"net/http"
	"net/url"
	"time"

	"go.uber.org/zap"
)

// ReservoirClient represents the client for the Reservoir API.
type ReservoirClient struct {
	Log *zap.SugaredLogger

	apiKey       string
	client       *http.Client
	baseURL      string
	requestDelay time.Duration
}

// NewReservoirClient creates a new Reservoir client with configuration.
func NewReservoirClient(apiKey string) *ReservoirClient {
	logger, _ := zap.NewProduction()

	return &ReservoirClient{
		Log: logger.Sugar(),

		apiKey:  apiKey,
		baseURL: "https://api.reservoir.tools",
		client: &http.Client{
			Timeout: time.Second * 10,
		},
		requestDelay: time.Millisecond * 250,
	}
}

// NewRequest creates a new request and adds authentication headers.
func (c *ReservoirClient) GetRequest(u *url.URL) *http.Request {
	req, _ := http.NewRequest("GET", u.String(), nil)

	req.Header.Set("X-API-KEY", c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	return req
}

// Get does a GET request.
func (c *ReservoirClient) Get(u *url.URL) (*http.Response, error) {
	req := c.GetRequest(u)
	return c.client.Do(req)
}
