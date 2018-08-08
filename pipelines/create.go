package pipelines

import (
	"fmt"
)

func (s *Service) Create(p Pipeline) error {
	var body interface{}
	err := s.client.Post(s.front50URL+"/pipelines", p, &body)
	if err != nil {
		return fmt.Errorf("could not create pipeline - %v", err)
	}
	return nil
}
