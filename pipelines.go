package plank

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

func (c *Client) pipelinesURL() string {
	return c.URLs["front50"] + "/pipelines"
}

// Get returns an array of all the Spinnaker pipelines
// configured for app
func (c *Client) GetPipelines(app string) ([]Pipeline, error) {
	var pipelines []Pipeline
	if err := c.Get(c.pipelinesURL()+"/"+app, &pipelines); err != nil {
		return nil, fmt.Errorf("could not get pipelines for %s - %v", app, err)
	}
	return pipelines, nil
}

// CreatePipeline creates a pipeline defined in the struct argument.
func (c *Client) CreatePipeline(p Pipeline) error {
	var unused interface{}
	if err := c.Post(c.pipelinesURL(), ApplicationJson, p, &unused); err != nil {
		return fmt.Errorf("could not create pipeline - %v", err)
	}
	return nil
}

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

// Execute a pipeline by application and pipeline.
func (c *Client) Execute(application, pipeline string) (*PipelineRef, error) {
	e := pipelineExecution{
		Enabled: true,
		Type:    "manual",
		DryRun:  false,
		User:    "anonymous",
	}
	var ref PipelineRef
	if err := c.Post(
		fmt.Sprintf("%s/%s/%s", c.pipelinesURL(), application, pipeline),
		ApplicationJson, e, &ref); err != nil {
		return nil, err
	}
	return &ref, nil
}
