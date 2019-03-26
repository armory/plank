package plank

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

// Application as returned from the Spinnaker API.
type Application struct {
	Name  string `json:"name" mapstructure:"name"`
	Email string `json:"email" mapstructure:"email"`
}

// Get returns the Application data struct for the
// given application name.
func (c *Client) GetApplication(name string, app *Application) error {
	err := c.Get(c.URLs["front50"]+"/v2/applications/"+name, app)
	return err
}

// Create an application.
func (c *Client) CreateApplication(a *Application) error {
	payload := newAppCreationRequest(a)
	var ref taskRefResponse
	err := c.Post(c.URLs["orca"]+"/ops", ApplicationContextJson, payload, &ref)
	if err != nil {
		return fmt.Errorf("could not create application - %v", err)
	}
	task, err := c.pollTaskStatus(ref.Ref)
	if err != nil || task.Status == "TERMINAL" {
		var errMsg string
		if err != nil {
			errMsg = err.Error()
		} else {
			errMsg = "received status TERMINAL"
		}
		return errors.New(fmt.Sprintf("failed to create application: %s", errMsg))
	}

	// This really shouldn't have to be here, but after the task to create an
	// app is marked complete sometimes the object still doesn't exist. So
	// after doing the create, and getting back a completion, we still need
	// to poll till we find the app in order to make sure future operations will
	// succeed.
	err = c.pollAppConfig(a.Name)
	return err
}

// todo: replace late night shortcut to not have to make all the structs.
const createAppTmpl = `{
	"application": "%s",
	"description": "Create Application: %s",
	"job" :[
		{
			"application": {
				"email": "%s",
				"name": "%s"
			},
			"type": "createApplication"
		}
	]
}`

func newAppCreationRequest(a *Application) map[string]interface{} {
	j := fmt.Sprintf(createAppTmpl, a.Name, a.Name, a.Email, a.Name)
	out := map[string]interface{}{}
	json.Unmarshal([]byte(j), &out)
	return out
}

// TODO:  All this task-based stuff should be pulled out into tasks.go and
// made re-usable.
type taskRefResponse struct {
	Ref string `json:"ref"`
}

type executionResponse struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	EndTime int    `json:"endTime"`
}

func (c *Client) pollTaskStatus(refURL string) (executionResponse, error) {
	if refURL == "" {
		return executionResponse{}, errors.New("no taskRef provided to follow")
	}
	timer := time.NewTimer(c.retryIncrement)
	t := time.NewTicker(1 * time.Second)
	defer t.Stop()

	for range t.C {
		var body executionResponse
		c.Get(c.URLs["orca"]+refURL, &body)
		if body.EndTime > 0 {
			return body, nil
		}
		select {
		case <-timer.C:
			return executionResponse{}, errors.New("timed out waiting for task to complete")
		default:
		}
	}
	return executionResponse{}, errors.New("exited poll loop before completion")
}

func (c *Client) getTask(refURL string) (executionResponse, error) {
	var body executionResponse
	err := c.Get(c.URLs["orca"]+refURL, &body)
	return body, err
}

func (c *Client) pollAppConfig(appName string) error {
	timer := time.NewTimer(c.retryIncrement)
	t := time.NewTicker(5 * time.Second)
	defer t.Stop()
	var app Application
	for range t.C {
		err := c.GetApplication(appName, &app)
		if err == nil {
			return nil
		}
		select {
		case <-timer.C:
			return fmt.Errorf("timed out waiting for app to appear - %v", err)
		default:
		}
	}
	return errors.New("exited poll loop before completion")
}
