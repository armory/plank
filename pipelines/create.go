package pipelines

// Create a pipeline using the DefaultService.
func Create(p Pipeline) error {
	return DefaultService.Create(p)
}

// Create a pipeline.
func (s *Service) Create(p Pipeline) error {
	return nil
}
