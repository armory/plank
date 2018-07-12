package fiat

// Fiat mirrors fiat.yml on disk
type Fiat struct {
	Auth struct {
		GroupMembership struct {
			Service string        `json:"service" mapstructure:"service"`
			Github  GithubAccount `json:"github" mapstructure:"github"`
		} `json:"groupMembership" mapstructure:"groupMembership"`
	}
	Services struct {
		Fiat struct {
			Enabled bool `json:"enabled" mapstructure:"enabled"`
		} `json:"fiat" mapstructure:"fiat"`
	} `json:"services" mapstructure:"services"`
}

// GithubAccount settings
type GithubAccount struct {
	Organization string `json:"organization" mapstructure:"organization"`
	BaseURL      string `json:"baseUrl" mapstructure:"baseUrl"`
	AccessToken  string `json:"access_token" mapstructure:"access_token"`
}
