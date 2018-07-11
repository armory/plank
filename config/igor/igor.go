package igor

type Igor struct {
	Jenkins struct {
		Enabled bool
		Masters JenkinsMasters
	}
	Travis struct {
		Enabled bool
		Masters TravisMasters
	}
}

type JenkinsMasters struct {
	Name     string
	Address  string
	Username string
	Password string
}

type TravisMasters struct {
	Name        string
	BaseURL     string
	Address     string
	GithubToken string
}
