package echo

// Echo mirrors the echo.yml file on disk
type Echo struct {
	Diagnostics struct {
		Enabled bool
		ID      string
	}
	Slack struct {
		Enabled bool
		Token   string
		botName string
	}
}
