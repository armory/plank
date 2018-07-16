package igor

// Igor mirrors igor.yml on disk
type Igor struct {
	Jenkins struct {
		Enabled bool           `json:"enabled,omitempty" mapstructure:"enabled"`
		Masters JenkinsMasters `json:"masters,omitempty" mapstructure:"masters"`
	} `json:"jenkins,omitempty" mapstructure:"jenkins"`
	Travis struct {
		Enabled bool          `json:"enabled,omitempty" mapstructure:"enabled"`
		Masters TravisMasters `json:"masters,omitempty" mapstructure:"masters"`
	}
}

// JenkinsMasters is Jenkins' masters settings
type JenkinsMasters struct {
	Name     string `json:"name,omitempty" mapstructure:"name"`
	Address  string `json:"address,omitempty" mapstructure:"address"`
	Username string `json:"username,omitempty" mapstructure:"username"`
	Password string `json:"password,omitempty" mapstructure:"password"`
}

// TravisMasters is Travis' masters settings
type TravisMasters struct {
	Name        string `json:"name,omitempty" mapstructure:"name"`
	BaseURL     string `json:"baseUrl,omitempty" mapstructure:"baseUrl"`
	Address     string `json:"address,omitempty" mapstructure:"address"`
	GithubToken string `json:"githubToken,omitempty" mapstructure:"githubToken"`
}
