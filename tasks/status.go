package tasks

// Status of a task using the DefaultService.
func Status(t ID) error {
	return DefaultService.Status(t)
}

// Status of a task.
func (s *Service) Status(t ID) error {
	return nil
}
