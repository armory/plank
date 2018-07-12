package kayenta

// Kayenta mirrors kayenta.yml on disk
type Kayenta struct {
	Kayenta struct {
		Datadog struct {
			Enabled  bool             `json:"enabled" mapstructure:"enabled"`
			Accounts []DatadogAccount `json:"accounts" mapstructure:"accounts"`
		} `json:"datadog" mapstructure:"datadog"`
	} `json:"kayenta" mapstructure:"kayenta"`
}

// DatadogAccount settings
type DatadogAccount struct {
	Name            string   `json:"name" mapstructure:"name"`
	APIKey          string   `json:"apiKey" mapstructure:"apiKey"`
	ApplicationKey  string   `json:"applicationKey" mapstructure:"applicationKey"`
	EndPointBaseURL string   `json:"endpoint.baseUrl" mapstructure:"endpoint.baseUrl"`
	SupportedTypes  []string `json:"supportedTypes" mapstructure:"supportedTypes"`
}
