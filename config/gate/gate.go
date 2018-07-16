package gate

type Gate struct {
	SAML SAML `json:"saml" mapstructure:"saml"`
	LDAP LDAP `json:"ldap" mapstructure:"ldap"`
	// https://docs.armory.io/install-guide/auth/#github
	// suggests github URIs used within standard OAuth implementation
	// TODO: OAuth structure should likely be used for: Github, Okta, Google, Facebook
	Security struct {
		OAuth2 OAuth2 `json:"oauth2" mapstructure:"oauth2"`
	} `json:"security" mapstructure:"security"`
	// TODO: Verify proper placement for OAuth
	// spinnaker docs put OAuth under spring:
	// armory production puts them under security:
	Spring struct {
		OAuth2 OAuth2 `json:"oauth2" mapstructure:"oauth2"`
	} `json:"spring" mapstructure:"spring"`
}

type SAML struct {
	Enabled  bool   `json:"enabled" mapstructure:"enabled"`
	IssuerID string `json:"issuerId" mapstructure:"issuerId"`
	Metadata string `json:"metadata" mapstructure:"metadata"`
	PEM      string `json:"pem" mapstructure:"pem"`
}

// LDAP setup https://docs.armory.io/install-guide/auth/#ldap-authentication
type LDAP struct {
	Enabled       bool   `json:"enabled" mapstructure:"enabled"`
	URL           string `json:"url" mapstructure:"url"`
	UserDNPattern string `json:"userDnPattern" mapstructure:"userDnPattern"`
}

// OAuth setup https://www.spinnaker.io/setup/security/authentication/oauth/
type OAuth2 struct {
	Client          OAuthClient    `json:"client" mapstructure:"client"`
	UserInfoMapping OAuthUIMapping `json:"userInfoMapping" mapstructure:"userInfoMapping"`
	Resource        struct {
		UserInfoURI string `json:"userInfoUri" mapstructure:"userInfoUri"`
	} `json:"resource" mapstructure:"resource"`
	ProviderRequirements struct {
		// Type is used to specifiy: github|google, etc
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
