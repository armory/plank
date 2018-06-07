package config

// Settings mirrors Spinnaker's yaml files.
type Settings struct {
	Services map[string]Service `json:"services" mapstructure:"services"`
}

// Service such as Fiat, Orca, Deck, Gate, etc.
type Service struct {
	Enabled bool   `json:"enabled" mapstructure:"enabled"`
	BaseURL string `json:"baseUrl" mapstructure:"baseUrl"`
}
