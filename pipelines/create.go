package pipelines

import (
	"fmt"
	"github.com/armory/plank/client"
)

func (s *Service) Create(p Pipeline) error {
	var body interface{}
	err := s.client.Post(s.front50URL+"/pipelines", client.ApplicationJson, p, &body)
	if err != nil {
		return fmt.Errorf("could not create pipeline - %v", err)
	}
	return nil
}
