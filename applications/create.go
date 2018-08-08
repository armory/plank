package applications

import (
	"github.com/armory/plank/tasks"
)

// CreationOption used when making a new application.
type CreationOption int

const (
	// WaitForCompletion of the Create operation before returning.
	WaitForCompletion CreationOption = iota
)

// Create an application using the DefaultService.
func Create(a Application, o ...CreationOption) (tasks.ID, error) {
	return DefaultService.Create(a)
}

// Create an application.
func (s *Service) Create(a Application, opts ...CreationOption) (tasks.ID, error) {
	for _, o := range opts {
		switch o {
		case WaitForCompletion:
			// block until complete
		}
	}
	return "", nil
}
