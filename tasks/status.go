package tasks

// TaskStatus of a task using the DefaultService.
func Status(t ID) (TaskStatus, error) {
	return Succeeded, nil
	//return DefaultService.TaskStatus(t)
}

// TaskStatus of a task.
func (s *Service) Status(t ID) error {
	return nil
}
