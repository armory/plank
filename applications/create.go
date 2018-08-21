package applications

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
	"github.com/armory/plank/client"
	"github.com/sirupsen/logrus"
)

// Create an application.
func (s *Service) Create(a Application) (Application, error) {
	payload := newAppCreationRequest(a)
	var ref taskRefResponse
	err := s.client.Post(s.orcaURL+"/ops", client.ApplicationContextJson, payload, &ref)
	if err != nil {
		return Application{}, fmt.Errorf("could not create application - %v", err)
	}
	logrus.Infof("Task creating application: '%s'", ref.Ref)
	task, err := s.pollTaskStatus(ref.Ref)
	if err != nil || task.Status == "TERMINAL" {
		var errMsg string
		if err != nil {
			errMsg = err.Error()
		} else {
			errMsg = "received status TERMINAL"
		}
		return Application{}, errors.New(fmt.Sprintf("failed to create application: %s", errMsg))
	}

	// This really shouldn't have to be here, but after the task to create an
	// app is marked complete sometimes the object still doesn't exist. So
	// after doing the create, and getting back a completion, we still need
	// to poll till we find the app in order to make sure future operations will
	// succeed.
	err = s.pollAppConfig(a.Name)
	return a, err
}

type taskRefResponse struct {
	Ref string `json:"ref"`
}

type executionResponse struct {
	ID      string `json:"id"`
	Status  string `json:"status"`
	EndTime int    `json:"endTime"`
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

func newAppCreationRequest(a Application) map[string]interface{} {
	j := fmt.Sprintf(createAppTmpl, a.Name, a.Name, a.Email, a.Name)
	out := map[string]interface{}{}
	json.Unmarshal([]byte(j), &out)
	return out
}

func (s *Service) pollTaskStatus(refURL string) (executionResponse, error) {
	if refURL == "" {
		return executionResponse{}, errors.New("no taskRef provided to follow")
	}
	timer := time.NewTimer(s.pollTime)
	t := time.NewTicker(1 * time.Second)
	defer t.Stop()

	for range t.C {
		var body executionResponse
		s.client.Get(s.orcaURL+refURL, &body)
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

func (s *Service) getTask(refURL string) (executionResponse, error) {
	var body executionResponse
	err := s.client.Get(s.orcaURL+refURL, &body)
	return body, err
}

func (s *Service) pollAppConfig(app string) error {
	timer := time.NewTimer(s.pollTime)
	t := time.NewTicker(5 * time.Second)
	defer t.Stop()
	for range t.C {
		_, err := s.Get(app)
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
