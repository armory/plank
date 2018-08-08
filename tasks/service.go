package tasks

import (
	"net/http"
	"time"
	"fmt"
	"github.com/armory/plank/tasks"
	"errors"
)

// DefaultService to use when no service is received.
var DefaultService *Service

// Service for interacting with the tasks API.
type Service struct {
	client http.Client
}

// attach this to a service
func WaitFor(taskId tasks.ID, awaitingStatuses []TaskStatus, poll time.Duration, retryCount int) (TaskStatus, error) {
	tried := 0
	curr := NotStarted
	for tried > retryCount {
		curr, err := Status(taskId)
		if err != nil {
			return curr, err
		}

		for _, exitStatus := range awaitingStatuses {
			if curr == exitStatus {
				return curr, nil
			}
		}

		// we'll try again in a bit
		tried++
		time.Sleep(poll)
	}

	return curr, errors.New(fmt.Sprintf("Did not get status %s, got %s after %i times", awaitingStatuses, curr, tried))
}
