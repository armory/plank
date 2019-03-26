package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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

type ContentType string

const (
	ApplicationJson        ContentType = "application/json"
	ApplicationContextJson ContentType = "application/context+json"
)

// Option for configuring a new Client.
type Option func(*Client) error

// New constructs a Client using the given http.Client-compatible client.
// Pass nil if you want to just use the default http.Client structure.
func New(httpClient *http.Client) (*Client, error) {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	c := &Client{
		http:           httpClient,
		retryIncrement: 100,
		maxRetry:       20,
	}
	return c, nil
}

// Get a JSON payload from the URL then decode it into the 'dest' arguement.
func (c *Client) Get(url string, dest interface{}) error {
	var reterr error
	for i := 0; i <= c.maxRetry; i++ {
		resp, err := c.http.Get(url)
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
func (c *Client) Post(url string, contentType ContentType, body interface{}, dest interface{}) error {
	var err error
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("could not post - body could not be marshaled to json - %v", err)
	}
	for i := 0; i <= c.maxRetry; i++ {
		resp, err := c.http.Post(url, string(contentType), bytes.NewBuffer(jsonBody))
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
		}

		// exponential back-off
		interval := c.retryIncrement * time.Duration(math.Exp2(float64(i)))
		time.Sleep(interval)
	}
	return err
}
