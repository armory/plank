package pipelines

import (
	"fmt"
)

// Pipeline is the structure that comes back from Spinnaker
// representing a pipeline definition (different than an execution)
type Pipeline struct {
	ID                   string                   `json:"id,omitempty"`
	Type                 string                   `json:"type,omitempty"`
	Name                 string                   `json:"name"`
	Application          string                   `json:"application"`
	Description          string                   `json:"description,omitempty"`
	ExecutionEngine      string                   `json:"executionEngine,omitempty"`
	Parallel             bool                     `json:"parallel"`
	LimitConcurrent      bool                     `json:"limitConcurrent"`
	KeepWaitingPipelines bool                     `json:"keepWaitingPipelines"`
	Stages               []map[string]interface{} `json:"stages,omitempty"`
	Triggers             []map[string]interface{} `json:"triggers,omitempty"`
	Parameters           []map[string]interface{} `json:"parameterConfig,omitempty"`
	Notifications        []map[string]interface{} `json:"notifications,omitempty"`
	LastModifiedBy       string                   `json:"lastModifiedBy"`
	Config               interface{}              `json:"config,omitempty"`
	UpdateTs             string                   `json:"updateTs"`
}

// Get returns an array of all the Spinnaker pipelines
// configured for app
func (s *Service) Get(app string) ([]Pipeline, error) {
	path := fmt.Sprintf(s.front50URL+"/pipelines/%s", app)
	var pipelines []Pipeline
	err := s.client.Get(path, &pipelines)
	if err != nil {
		return nil, fmt.Errorf("could not get pipelines for %s - %v", app, err)
	}
	return pipelines, nil
}
