package applications

import (
	"github.com/armory/plank/tasks"
)

// Create an application using the DefaultService.
func Create(a Application) (tasks.ID, error) {
	return DefaultService.Create(a)
}

// Create an application.
func (s *Service) Create(a Application) (tasks.ID, error) {
	return "", nil
}
