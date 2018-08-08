package applications

import (
	"github.com/armory/plank/tasks"
	"net/http"
	log "github.com/sirupsen/logrus"
	"encoding/json"
	"time"
	"fmt"
)

// CreationOption used when making a new application.
type CreationOption int

const (
	// WaitForCompletion of the Create operation before returning.
	WaitForCompletion CreationOption = iota
)

// Create an application using the DefaultService.
func Create(appCreateParams Application, opts ...CreationOption) (Application, error) {
	var app Application

	// this should be simplier by calling the service create
	for _, o := range opts {
		switch o {
		case WaitForCompletion:
			taskId, err := DefaultService.Create(appCreateParams, opts...)
			if err != nil {
				return app, err
			}

			// check the status before returning the app
			tasks.WaitFor(taskId, []tasks.TaskStatus{tasks.Succeeded, tasks.Terminal}, time.Second*2, 10)

			return DefaultService.Get(app.Name)
		default:
			_, err := DefaultService.Create(appCreateParams, opts...)
			return app, err
		}
	}
}

// Create an application.
func (s *Service) Create(a Application, opts ...CreationOption) (tasks.ID, error) {
	for _, o := range opts {
		switch o {
		case WaitForCompletion:
			client := &http.Client{}
			res, _ := client.Get("http://spinnaker.dev.armory.io:8084/applications/armoryhellodeploy/tasks/01CK71Q248ZNFNBHZMEGCAYEK7")
			decoder := json.NewDecoder(res.Body)

			task := tasks.Task{}
			decoder.Decode(&task)
			log.Info(task)

			// close it here because we're doing it in a loop
			res.Body.Close()
		default:
			return nil, nil
		}
	}
}

func Get(name string) (Application, error) {
	return DefaultService.Get(name)
}

func (s *Service) Get(name string) (Application, error) {
	client := &http.Client{}
	res, _ := client.Get(fmt.Sprintf("http://spinnaker.dev.armory.io:8084/applications/%s", name))
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)

	var app Application
	decoder.Decode(&app)
	log.Info(app)
}
