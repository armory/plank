package client

import (
	"encoding/json"
	"math"
	"net/http"
	"time"
)

// Client for working with API servers that accept and return JSON payloads.
type Client struct {
	http           *http.Client
	retryIncrement time.Duration
	maxRetry       int
}

func New(options ...func(*Client)) *Client {
	c := &Client{
		http:           &http.Client{},
		retryIncrement: 100,
		maxRetry:       20,
	}
	for _, option := range options {
		option(c)
	}
	return c
}

// Get a JSON payload from the URL then decode it into the 'dest' arguement.
func (c *Client) Get(url string, dest interface{}) error {
	var err error
	for i := 0; i <= c.maxRetry; i++ {
		resp, err := c.http.Get(url)
		success := resp.StatusCode >= 200 && resp.StatusCode < 400
		if success && err == nil {
			err := json.NewDecoder(resp.Body).Decode(&dest)
			defer resp.Body.Close()
			if err != nil {
				return err
			}
			return nil
		}
		// exponential back-off
		interval := c.retryIncrement * time.Duration(math.Exp2(float64(i)))
		time.Sleep(interval)
	}
	return err
}
