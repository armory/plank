// Package permissions helps determine permissions for users.
package permissions

import (
	"github.com/armory/plank/client"
	"github.com/armory/plank/config"
)

const (
	defaultFiatBaseURL = "http://fiat"
)

type getter interface {
	Get(path string, dest interface{}) error
}

// Service for interacting with the permissions API.
type Service struct {
	client getter
}

// Option for configuring a service.
type Option func(*Service) error

// NewService for checking permissions.
func NewService(options ...Option) *Service {
	defClient, _ := client.New(client.BaseURL(defaultFiatBaseURL))
	s := &Service{
		client: defClient,
	}
	for _, option := range options {
		// TODO: handle errors
		option(s)
	}
	return s
}

// Client option for a new permissions service.
func Client(c *client.Client) Option {
	return func(s *Service) error {
		s.client = c
		// TODO: validation
		return nil
	}
}

// Settings based configuration option for a new permissions service.
func Settings(fiat *config.Service) Option {
	return func(s *Service) error {
		c, err := client.New(client.BaseURL(fiat.BaseURL))
		if err != nil {
			return err
		}
		s.client = c
		return nil
	}
}
