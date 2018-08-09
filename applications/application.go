package applications

// Application as returned from the Spinnaker API.
type Application struct {
	Name  string `json:"name" mapstructure:"name"`
	Email string `json:"email" mapstructure:"email"`
}

// Get returns the Application data struct for the
// given application name.
func (s *Service) Get(name string) (Application, error) {
	var app Application
	err := s.client.Get(s.orcaURL+"/v2/applications/"+name, &app)
	return app, err
}
