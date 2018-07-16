package gate

type Gate struct {
	SAML SAML `json:"saml,omitempty" mapstructure:"saml"`
	LDAP LDAP `json:"ldap,omitempty" mapstructure:"ldap"`
	// https://docs.armory.io/install-guide/auth/#github
	// suggests github URIs used within standard OAuth implementation
	// TODO: OAuth structure should likely be used for: Github, Google, Facebook
	Security struct {
		OAuth2 OAuth2 `json:"oauth2,omitempty" mapstructure:"oauth2"`
	} `json:"security,omitempty" mapstructure:"security"`
}

type SAML struct {
	Enabled  bool   `json:"enabled,omitempty" mapstructure:"enabled"`
	IssuerID string `json:"issuerId,omitempty" mapstructure:"issuerId"`
	Metadata string `json:"metadata,omitempty" mapstructure:"metadata"`
	PEM      string `json:"pem,omitempty" mapstructure:"pem"`
}

// LDAP setup https://docs.armory.io/install-guide/auth/#ldap-authentication
type LDAP struct {
	Enabled       bool   `json:"enabled,omitempty" mapstructure:"enabled"`
	URL           string `json:"url,omitempty" mapstructure:"url"`
	UserDNPattern string `json:"userDnPattern,omitempty" mapstructure:"userDnPattern"`
}

// OAuth setup https://www.spinnaker.io/setup/security/authentication/oauth/
type OAuth2 struct {
	Client          OAuthClient    `json:"client,omitempty" mapstructure:"client"`
	UserInfoMapping OAuthUIMapping `json:"userInfoMapping,omitempty" mapstructure:"userInfoMapping"`
	Resource        struct {
		UserInfoURI string `json:"userInfoUri,omitempty" mapstructure:"userInfoUri"`
	} `json:"resource,omitempty" mapstructure:"resource"`
	ProviderRequirements struct {
		// Type is used to specifiy: github|google, etc
		Type         string `json:"requirementsType,omitempty" mapstructure:"requirementsType"`
		Organization string `json:"requirementsOrg,omitempty" mapstructure:"requirementsOrg"`
	} `json:"providerRequirements,omitempty" mapstructure:"providerRequirements"`
}

type OAuthClient struct {
	ClientID             string `json:"clientId,omitempty" mapstructure:"clientId"`
	ClientSecret         string `json:"clientSecret,omitempty" mapstructure:"clientSecret"`
	UserAuthorizationURI string `json:"userAuthorizationUri,omitempty" mapstructure:"userAuthorizationUri"`
	AccessTokenURI       string `json:"accessTokenUri,omitempty" mapstructure:"accessTokenUri"`
	Scope                string `json:"scope,omitempty" mapstructure:"scope"`
}

type OAuthUIMapping struct {
	Email     string `json:"email,omitempty" mapstructure:"email"`
	FirstName string `json:"firstName,omitempty" mapstructure:"firstName"`
	LastName  string `json:"lastName,omitempty" mapstructure:"lastName"`
	Username  string `json:"username,omitempty" mapstructure:"username"`
}
