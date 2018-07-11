package config

// Dinghy mirrors dinghy.yml files on disk
type Dinghy struct {
	TemplateOrg       string  `json:"templateOrg" mapstructure:"templateOrg"`
	TemplateRepo      string  `json:"templateRepo" mapstructure:"templateRepo"`
	AutoLockPipelines bool    `json:"autoLockPipelines" mapstructure:"autoLockPipelines"`
	CertPath          string  `json:"certPath" mapstructure:"certPath"`
	GithubCredsPath   string  `json:"githubCredsPath" mapstructure:"githubCredsPath"`
	Filename          string  `json:"dinghyFilename" mapstructure:"dinghyFilename"`
	SpinAPIURL        string  `json:"spinAPIUrl" mapstructure:"spinAPIUrl"`
	SpinUIURL         string  `json:"spinUIUrl" mapstructure:"spinUIUrl"`
	Orca              Orca    `json:"orca" mapstructure:"orca"`
	Fiat              Fiat    `json:"fiat" mapstructure:"fiat"`
	Logging           Logging `json:"logging" mapstructure:"logging"`
	Front50           struct {
		Service
	} `json:"front50" mapstructure:"front50"`
}

// Orca settings
type Orca struct {
	Service
}

// Fiat settings
type Fiat struct {
	Service
	AuthUser string `json:"authUser" mapstructure:"authUser"`
}

// Logging settings
type Logging struct {
	Level string `json:"level" mapstructure:"level"`
}
