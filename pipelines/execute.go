package pipelines

// Execute a pipeline using the DefaultService.
func Execute(id ID) error {
	return DefaultService.Execute(id)
}

// Execute a pipeline.
func (s *Service) Execute(p ID) error {
	return nil
}
