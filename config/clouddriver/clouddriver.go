package clouddriver

// Clouddriver mirrors clouddriver.yaml files on disk
type Clouddriver struct {
	DockerRegistry struct {
		Enabled  bool            `json:"enabled" mapstructure:"enabled"`
		Accounts []DockerAccount `json:"accounts" mapstructure:"accounts"`
	} `json:"dockerRegistry" mapstructure:"dockerRegistry"`

	Kubernetes struct {
		Enabled  bool                `json:"enabled" mapstructure:"enabled"`
		Accounts []KubernetesAccount `json:"accounts" mapstructure:"accounts"`
	} `json:"kubernetes" mapstructure:"kubernetes"`

	AWS struct {
		Accounts []AWSAccount `json:"accounts" mapstructure:"accounts"`
	} `json:"aws" mapstructure:"aws"`

	GCP struct {
		Enabled  bool         `json:"enabled" mapstructure:"enabled"`
		Accounts []GCPAccount `json:"accounts" mapstructure:"accounts"`
	} `json:"google" mapstructure:"google"`

	Artifacts struct {
		Github struct {
			Enabled  bool                    `json:"enabled" mapstructure:"enabled"`
			Accounts []GithubArtifactAccount `json:"accounts" mapstructure:"accounts"`
		} `json:"github" mapstructure:"github"`
	} `json:"artifacts" mapstructure:"artifacts"`
}

// GithubArtifactAccount settings
type GithubArtifactAccount struct {
	Name     string `json:"name" mapstructure:"name"`
	Username string `json:"username" mapstructure:"username"`
	Token    string `json:"token" mapstructure:"token"`
}

// DockerAccount settings
type DockerAccount struct {
	Name         string   `json:"name" mapstructure:"name"`
	Username     string   `json:"username" mapstructure:"username"`
	PasswordFile string   `json:"passwordFile" mapstructure:"passwordFile"`
	Address      string   `json:"address" mapstructure:"address"`
	Repositories []string `json:"repositories" mapstructure:"repositories"`
}

// KubernetesAccount settings
type KubernetesAccount struct {
	Name             string           `json:"name" mapstructure:"name"`
	Namespaces       []string         `json:"namespaces" mapstructure:"namespaces"`
	KubeconfigFile   string           `json:"kubeconfigFile" mapstructure:"kubeconfigFile"`
	ProviderVersion  string           `json:"providerVersion" mapstructure:"providerVersion"`
	DockerRegistries []DockerRegistry `json:"dockerRegistries" mapstructure:"dockerRegistries"`
}

// DockerRegistry settings
type DockerRegistry struct {
	AccountName string `json:"accountName" mapstructure:"accountName"`
}

// AWSAccount settings
type AWSAccount struct {
	Name      string      `json:"name" mapstructure:"name"`
	AccountID string      `json:"accountId" mapstructure:"accountId"`
	Regions   []AWSRegion `json:"regions" mapstructure:"regions"`
}

// AWSRegion settings
type AWSRegion struct {
	Name string `json:"name" mapstructure:"name"`
}

// GCPAccount settings
type GCPAccount struct {
	Name     string `json:"name" mapstructure:"name"`
	Project  string `json:"project" mapstructure:"project"`
	JSONPath string `json:"jsonPath" mapstructure:"jsonPath"`
}
