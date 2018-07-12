package igor

// Igor mirrors igor.yml on disk
type Igor struct {
	Jenkins struct {
		Enabled bool           `json:"enabled" mapstructure:"enabled"`
		Masters JenkinsMasters `json:"masters" mapstructure:"masters"`
	} `json:"jenkins" mapstructure:"jenkins"`
	Travis struct {
		Enabled bool          `json:"enabled" mapstructure:"enabled"`
		Masters TravisMasters `json:"masters" mapstructure:"masters"`
	}
}

// JenkinsMasters is Jenkins' masters settings
type JenkinsMasters struct {
	Name     string `json:"name" mapstructure:"name"`
	Address  string `json:"address" mapstructure:"address"`
	Username string `json:"username" mapstructure:"username"`
	Password string `json:"password" mapstructure:"password"`
}

// TravisMasters is Travis' masters settings
type TravisMasters struct {
	Name        string `json:"name" mapstructure:"name"`
	BaseURL     string `json:"baseUrl" mapstructure:"baseUrl"`
	Address     string `json:"address" mapstructure:"address"`
	GithubToken string `json:"githubToken" mapstructure:"githubToken"`
}
