package fiat

// Fiat mirrors fiat.yml on disk
type Fiat struct {
	Auth struct {
		GroupMembership struct {
			Service string        `json:"service,omitempty" mapstructure:"service"`
			Github  GithubAccount `json:"github,omitempty" mapstructure:"github"`
		} `json:"groupMembership,omitempty" mapstructure:"groupMembership"`
	}
	Services struct {
		Fiat struct {
			Enabled bool `json:"enabled,omitempty" mapstructure:"enabled"`
		} `json:"fiat,omitempty" mapstructure:"fiat"`
	} `json:"services,omitempty" mapstructure:"services"`
}

// GithubAccount settings
type GithubAccount struct {
	Organization string `json:"organization,omitempty" mapstructure:"organization"`
	BaseURL      string `json:"baseUrl,omitempty" mapstructure:"baseUrl"`
	AccessToken  string `json:"access_token,omitempty" mapstructure:"access_token"`
}
