package pipelines

import (
	"github.com/armory/plank/tasks"
)

// Create a pipeline using the DefaultService.
func Create(p Pipeline) (tasks.ID, error) {
	return DefaultService.Create(p)
}

// Create a pipeline.
func (s *Service) Create(p Pipeline) (tasks.ID, error) {
	return "", nil
}
