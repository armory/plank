package echo

// Echo mirrors the echo.yml file on disk
type Echo struct {
	Diagnostics struct {
		Enabled bool   `json:"enabled,omitempty" mapstructure:"enabled"`
		ID      string `json:"id,omitempty" mapstructure:"id"`
	} `json:"diagnostics,omitempty" mapstructure:"diagnostics"`

	Slack struct {
		Enabled bool   `json:"enabled,omitempty" mapstructure:"enabled"`
		Token   string `json:"token,omitempty" mapstructure:"token"`
		BotName string `json:"botName,omitempty" mapstructure:"botName"`
	} `json:"slack,omitempty" mapstructure:"slack"`

	ArmoryWebHooks struct {
		Enabled    bool `json:"enabled,omitempty" mapstructure:"enabled"`
		Forwarding struct {
			BaseURL  string `json:"baseUrl,omitempty" mapstructure:"baseUrl"`
			EndPoint string `json:"endpoint,omitempty" mapstructure:"endpoint"`
		} `json:"forwarding,omitempty" mapstructure:"forwarding"`
	} `json:"armorywebhooks,omitempty" mapstructure:"armorywebhooks"`
}
