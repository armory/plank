package pipelines

// Execute a pipeline using the DefaultService.
func Execute(p ID) error {
	return DefaultService.Execute(p)
}

// Execute a pipeline.
func (s *Service) Execute(p ID) error {
	return nil
}
