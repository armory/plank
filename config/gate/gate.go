package gate

type Gate struct {
	Saml   Saml   `json:"saml" mapstructure:"saml"`
	Okta   Okta   `json:"okta" mapstructure:"okta"`
	LDAP   LDAP   `json:"ldap" mapstructure:"ldap"`
	Github Github `json:"github" mapstructure:"github"`
	// TODO: Verify proper placement for OAuth
	// spinnaker docs put OAuth under spring:
	// armory production puts them under security:
	Spring struct {
		OAuth OAuth `json:"oauth" mapstructure:"oauth"`
	} `json:"spring" mapstructure:"spring"`
}

type Saml struct {
	Enabled  bool   `json:"enabled" mapstructure:"enabled"`
	IssuerID string `json:"issuerId" mapstructure:"issuerId"`
	Metadata string `json:"metadata" mapstructure:"metadata"`
	Pem      string `json:"pem" mapstructure:"pem"`
}

type Okta struct {
	Enabled  bool   `json:"enabled" mapstructure:"enabled"`
	IssuerID string `json:"issuerId" mapstructure:"issuerId"`
	Metadata string `json:"metadata" mapstructure:"metadata"`
	Pem      string `json:"pem" mapstructure:"pem"`
}

// LDAP setup https://docs.armory.io/install-guide/auth/#ldap-authentication
type LDAP struct {
	Enabled       bool   `json:"enabled" mapstructure:"enabled"`
	URL           string `json:"url" mapstructure:"url"`
	userDNPattern string `json:"userDnPattern" mapstructure:"userDnPattern"`
}

type Github struct {
}

// OAuth setup https://www.spinnaker.io/setup/security/authentication/oauth/
type OAuth struct {
	Client          OAuthClient    `json:"client" mapstructure:"client"`
	UserInfoMapping OAuthUIMapping `json:"userInfoMapping" mapstructure:"userInfoMapping"`
	Resource        struct {
		UserInfoURI string `json:"userInfoUri" mapstructure:"userInfoUri"`
	} `json:"resource" mapstructure:"resource"`
	ProviderRequirements struct {
		Type         string `json:"requirementsType" mapstructure:"requirementsType"`
		Organization string `json:"requirementsOrg" mapstructure:"requirementsOrg"`
	} `json:"providerRequirements" mapstructure:"providerRequirements"`
}

type OAuthClient struct {
	ClientID             string `json:"clientId" mapstructure:"clientId"`
	ClientSecret         string `json:"clientSecret" mapstructure:"clientSecret"`
	UserAuthorizationURI string `json:"userAuthorizationUri" mapstructure:"userAuthorizationUri"`
	AccessTokenURI       string `json:"accessTokenUri" mapstructure:"accessTokenUri"`
	Scope                string `json:"scope" mapstructure:"scope"`
}

type OAuthUIMapping struct {
	Email     string `json:"email" mapstructure:"email"`
	FirstName string `json:"firstName" mapstructure:"firstName"`
	LastName  string `json:"lastName" mapstructure:"lastName"`
	Username  string `json:"username" mapstructure:"username"`
}
