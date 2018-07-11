package fiat

type Fiat struct {
	Auth struct {
		GroupMembership struct {
			Service string
			Github  GithubAccount
		}
	}
	Services struct {
		Fiat struct {
			Enabled bool
		}
	}
}

type GithubAccount struct {
	Organization string
	BaseURL      string
	AccessToken  string `json:"access_token"`
}
