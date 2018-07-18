package spinnaker

// Spinnaker mirrors spinnaker.yaml files on disk
type Spinnaker struct {
	Services Services `json:"services,omitempty" mapstructure:"services"`
	Logging  Logging  `json:"logging,omitempty" mapstructure:"logging"`
}

// Services within Spinnaker.
type Services struct {
	Fiat     Service `json:"fiat,omitempty" mapstructure:"fiat"`
	Front50  Front50 `json:"front50,omitempty" mapstructure:"front50"`
	Jenkins  Jenkins `json:"jenkins,omitempty" mapstructure:"jenkins"`
	Redis    Redis   `json:"redis,omitempty" mapstructure:"redis"`
	Features struct {
		Jira Jira `json:"jira,omitempty" mapstructure:"jira"`
	} `json:"features,omitempty" mapstructure:"features"`
	Deck struct {
		Hostname string `json:"hostname,omitempty" mapstructure:"hostname"`
	} `json:"deck,omitempty" mapstructure:"deck"`
}

type Redis struct {
	Host string `json:"host,omitempty" mapstructure:"host"`
	Port string `json:"port,omitempty" mapstructure:"port"`
}

type Jira struct {
	Enabled        bool   `json:"enabled,omitempty" mapstructure:"enabled"`
	BasicAuthToken string `json:"basicAuthToken,omitempty" mapstructure:"basicAuthToken"`
	CreateIssueURL string `json:"createIssueUrl,omitempty" mapstructure:"createIssueUrl"`
}

// Jenkins service settings
type Jenkins struct {
	Enabled       bool `json:"enabled,omitempty" mapstructure:"enabled"`
	DefaultMaster struct {
		Name     string `json:"name,omitempty" mapstructure:"name"`
		BaseURL  string `json:"baseUrl,omitempty" mapstructure:"baseUrl"`
		Username string `json:"username,omitempty" mapstructure:"username"`
		Password string `json:"password,omitempty" mapstructure:"password"`
	} `json:"defaultMaster,omitempty" mapstructure:"defaultMaster"`
}

// Front50 service settings.
type Front50 struct {
	Service
	StorageBucket string `json:"storage_bucket,omitempty" mapstructure:"storage_bucket"`
	StoragePrefix string `json:"rootFolder,omitempty" mapstructure:"rootFolder"`
	S3            struct {
		Enabled bool `json:"enabled,omitempty" mapstructure:"enabled"`
	} `json:"s3,omitempty" mapstructure:"s3"`
	GCS struct {
		Enabled bool `json:"enabled,omitempty" mapstructure:"enabled"`
	} `json:"gcs,omitempty" mapstructure:"gcs"`
	Redis struct {
		Enabled bool `json:"enabled,omitempty" mapstructure:"enabled"`
	} `json:"redis,omitempty" mapstructure:"redis"`
}

// Service such as Fiat, Orca, Deck, Gate, etc.
type Service struct {
	Enabled bool   `json:"enabled,omitempty" mapstructure:"enabled"`
	BaseURL string `json:"baseUrl,omitempty" mapstructure:"baseUrl"`
}

// Logging levels.
type Logging struct {
	Level struct {
		Spinnaker string `json:"com.netflix.spinnaker,omitempty" mapstructure:"com.netflix.spinnaker"`
		Root      string `json:"root,omitempty" mapstructure:"root"`
	} `json:"level,omitempty" mapstructure:"level"`
}
