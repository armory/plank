package pipelines

import (
	"github.com/armory/plank/client"
)

// DefaultService to use when no service is received.
var DefaultService *Service

var (
	defaultFront50URL = "http://armory-front50:8080"
	defaultGateURL    = "http://armory-gate:8084"
)

type postGetter interface {
	Post(path string, body, dest interface{}) error
	Get(path string, dest interface{}) error
}

// Service for interacting with the applications API.
type Service struct {
	client     postGetter
	front50URL string
	gateURL    string
}

// Option for configuring a service.
type Option func(*Service) error

// NewService for checking permissions.
func NewService(options ...Option) *Service {
	defClient, _ := client.New()
	s := &Service{
		client:     defClient,
		front50URL: defaultFront50URL,
	}
	for _, option := range options {
		// TODO: handle errors
		option(s)
	}
	return s
}

// Front50URL option to change the URL used to talk to Orca.
func Front50URL(url string) Option {
	return func(s *Service) error {
		s.front50URL = url
		return nil
	}
}

// GateURL option to change the URL used to talk to Orca.
func GateURL(url string) Option {
	return func(s *Service) error {
		s.gateURL = url
		return nil
	}
}
