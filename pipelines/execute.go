package pipelines

import (
	"net/url"
	"github.com/armory/plank/client"
)

type pipelineExecution struct {
	Enabled bool   `json:"enabled"`
	Type    string `json:"type"`
	DryRun  bool   `json:"dryRun"`
	User    string `json:"user"`
}

type PipelineRef struct {
	// Ref is the path the the execution. Use it to get status updates.
	Ref string `json:"ref"`
}

// Execute a pipeline by application name and pipeline name.
func (s *Service) Execute(application, pipelineName string) (PipelineRef, error) {
	u, _ := url.Parse(s.gateURL + "/pipelines/" + application + "/" + pipelineName)
	e := pipelineExecution{
		Enabled: true,
		Type:    "manual",
		DryRun:  false,
		User:    "anonymous",
	}
	var ref PipelineRef
	err := s.client.Post(u.String(), client.ApplicationJson, e, &ref)
	return ref, err
}
