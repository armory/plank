package client

import (
	"encoding/json"
	"errors"
	"math"
	"net/http"
	"time"
)

// Client for working with API servers that accept and return JSON payloads.
type Client struct {
	baseURL        string
	http           *http.Client
	retryIncrement time.Duration
	maxRetry       int
}

// Option for configuring a new Client.
type Option func(*Client) error

// New constructs a Client using functional based option arguements.
//
// You can read more about the merits of this approach here:
//   https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis
// and here:
//   https://github.com/tmrts/go-patterns/blob/master/idiom/functional-options.md
func New(options ...Option) (*Client, error) {
	c := &Client{
		http:           &http.Client{},
		retryIncrement: 100,
		maxRetry:       20,
	}
	for _, option := range options {
		err := option(c)
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}

// BaseURL for the client to use.
func BaseURL(address string) Option {
	return func(c *Client) error {
		if len(address) == 0 {
			return errors.New("baseUrl can not be empty")
		}
		c.baseURL = address
		if c.baseURL[len(c.baseURL)-1] == '/' {
			c.baseURL = c.baseURL[:len(c.baseURL)-1]
		}
		// TODO: validation
		return nil
	}
}

// Get a JSON payload from the URL then decode it into the 'dest' arguement.
func (c *Client) Get(path string, dest interface{}) error {
	if len(path) == 0 || path[0] != '/' {
		return errors.New("invalid path. must start with '/'")
	}
	var err error
	for i := 0; i <= c.maxRetry; i++ {
		resp, err := c.http.Get(c.baseURL + path)
		success := resp.StatusCode >= 200 && resp.StatusCode < 400
		if success && err == nil {
			err := json.NewDecoder(resp.Body).Decode(dest)
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
