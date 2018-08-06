package pipelines

// Status of a pipeline using the DefaultService.
func Status(p ID) error {
	return DefaultService.Status(p)
}

// Status of a pipeline.
func (s *Service) Status(p ID) error {
	return nil
}
