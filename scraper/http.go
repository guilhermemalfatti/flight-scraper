package scraper

import (
	"net/http"
	"time"
)

const (
	DefaultHTTPTimeout = time.Second * 10
)

type ClientConfig struct {
	Timeout time.Duration
}

func (c *ClientConfig) TimeoutOrDefault() time.Duration {
	if c.Timeout == 0 {
		return DefaultHTTPTimeout
	}

	return c.Timeout
}

func NewHttpClient(config *ClientConfig) *http.Client {
	if config == nil {
		config = &ClientConfig{}
	}

	client := &http.Client{
		Timeout: config.TimeoutOrDefault(),
	}

	return client
}
