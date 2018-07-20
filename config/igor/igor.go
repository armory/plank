package igor

// Igor mirrors igor.yml on disk
type Igor struct {
	Jenkins struct {
		Enabled bool            `json:"enabled,omitempty" mapstructure:"enabled"`
		Masters []JenkinsMaster `json:"masters,omitempty" mapstructure:"masters"`
	} `json:"jenkins,omitempty" mapstructure:"jenkins"`
	Travis struct {
		Enabled bool           `json:"enabled,omitempty" mapstructure:"enabled"`
		Masters []TravisMaster `json:"masters,omitempty" mapstructure:"masters"`
	} `json:"travis,omitempty" mapstructure:"travis"`
}

// JenkinsMaster is Jenkins' masters settings
type JenkinsMaster struct {
	Name     string `json:"name,omitempty" mapstructure:"name"`
	Address  string `json:"address,omitempty" mapstructure:"address"`
	Username string `json:"username,omitempty" mapstructure:"username"`
	Password string `json:"password,omitempty" mapstructure:"password"`
	CSRF     bool   `json:"csrf,omitempty" mapstructure:"csrf"`
}

// TravisMaster is Travis' masters settings
type TravisMaster struct {
	Name        string `json:"name,omitempty" mapstructure:"name"`
	BaseURL     string `json:"baseUrl,omitempty" mapstructure:"baseUrl"`
	Address     string `json:"address,omitempty" mapstructure:"address"`
	GithubToken string `json:"githubToken,omitempty" mapstructure:"githubToken"`
}
