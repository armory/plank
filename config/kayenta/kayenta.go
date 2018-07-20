package kayenta

// Kayenta mirrors kayenta.yml on disk
type Kayenta struct {
	Kayenta struct {
		Datadog struct {
			Enabled  bool             `json:"enabled,omitempty" mapstructure:"enabled"`
			Accounts []DatadogAccount `json:"accounts,omitempty" mapstructure:"accounts"`
		} `json:"datadog,omitempty" mapstructure:"datadog"`
	} `json:"kayenta,omitempty" mapstructure:"kayenta"`
}

// DatadogAccount settings
type DatadogAccount struct {
	Name            string   `json:"name,omitempty" mapstructure:"name"`
	APIKey          string   `json:"apiKey,omitempty" mapstructure:"apiKey"`
	ApplicationKey  string   `json:"applicationKey,omitempty" mapstructure:"applicationKey"`
	EndPointBaseURL string   `json:"endpoint.baseUrl,omitempty" mapstructure:"endpoint.baseUrl"`
	SupportedTypes  []string `json:"supportedTypes,omitempty" mapstructure:"supportedTypes"`
}
