package config

// Settings mirrors Spinnaker's yaml files.
type Settings struct {
	Services Services `json:"services" mapstructure:"services"`
}

// Services within Spinnaker.
type Services struct {
	Fiat    Service `json:"fiat" mapstructure:"fiat"`
	Front50 Front50 `json:"front50" mapstructure:"front50"`
}

// Front50 service settings.
type Front50 struct {
	Service
	StorageBucket string `json:"storage_bucket" mapstructure:"storage_bucket"`
	S3            struct {
		Enabled bool `json:"enabled" mapstructure:"enabled"`
	} `json:"s3" mapstructure:"s3"`
	GCS struct {
		Enabled bool `json:"enabled" mapstructure:"enabled"`
	} `json:"gcs" mapstructure:"gcs"`
}

// Service such as Fiat, Orca, Deck, Gate, etc.
type Service struct {
	Enabled bool   `json:"enabled" mapstructure:"enabled"`
	BaseURL string `json:"baseUrl" mapstructure:"baseUrl"`
}
