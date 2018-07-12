package gate

type Gate struct {
	Saml Saml `json:"saml" mapstructure:"saml"`
}

type Saml struct {
	Enabled  bool   `json:"enabled" mapstructure:"enabled"`
	IssuerID string `json:"issuerId" mapstructure:"issuerId"`
	Metadata string `json:"metadata" mapstructure:"metadata"`
	Pem      string `json:"pem" mapstructure:"pem"`
}
