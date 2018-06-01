package settings

// S holds default configuration used by the package.
type S struct {
	Services struct {
		Fiat service `json:"fiat"`
	} `json:"services"`
}

type service struct {
	BaseURL string `json:"baseUrl"`
}
