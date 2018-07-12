package echo

// Echo mirrors the echo.yml file on disk
type Echo struct {
	Diagnostics struct {
		Enabled bool   `json:"enabled" mapstructure:"enabled"`
		ID      string `json:"id" mapstructure:"id"`
	} `json:"diagnostics" mapstructure:"diagnostics"`

	Slack struct {
		Enabled bool   `json:"enabled" mapstructure:"enabled"`
		Token   string `json:"token" mapstructure:"token"`
		BotName string `json:"botName" mapstructure:"botName"`
	} `json:"slack" mapstructure:"slack"`
}
