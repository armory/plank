package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"
	"time"
	"io/ioutil"
)

// Client for working with API servers that accept and return JSON payloads.
type Client struct {
	baseURL        string
	http           *http.Client
	retryIncrement time.Duration
	maxRetry       int
}

type ContentType string

const (
	ApplicationJson        ContentType = "application/json"
	ApplicationContextJson ContentType = "application/context+json"
)

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

// MaxRetry set to a number.
func MaxRetry(num int) Option {
	return func(c *Client) error {
		c.maxRetry = num
		return nil
	}
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
	var reterr error
	for i := 0; i <= c.maxRetry; i++ {
		resp, err := c.http.Get(c.url(path))
		if resp == nil {
			reterr = errors.New("empty response")
		}
		if err == nil && resp.StatusCode >= 200 && resp.StatusCode < 400 {
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
	return reterr
}

// Post a JSON payload from the URL then decode it into the 'dest' arguement.
func (c *Client) Post(path string, contentType ContentType, body interface{}, dest interface{}) error {
	var err error
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("could not post - body could not be marshaled to json - %v", err)
	}
	for i := 0; i <= c.maxRetry; i++ {
		resp, err := c.http.Post(c.url(path), string(contentType), bytes.NewBuffer(jsonBody))
		if resp == nil {
			err = errors.New("empty response")
		}

		if err == nil && resp.StatusCode >= 200 && resp.StatusCode < 400 {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
			}

			if len(b) > 0 {
				err := json.Unmarshal(b, &dest)
				if err != nil {
					return err
				}
				return nil
			}
			continue // something went funky, let's try again
		}

		// exponential back-off
		interval := c.retryIncrement * time.Duration(math.Exp2(float64(i)))
		time.Sleep(interval)
	}
	return err
}

func (c *Client) url(path string) string {
	if c.baseURL != "" {
		return c.baseURL + path
	}
	return path
}
