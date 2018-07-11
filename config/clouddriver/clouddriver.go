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
	} `json:"Kubernetes" mapstructure:"Kubernetes"`

	AWS struct {
		Accounts []AWSAccount `json:"accounts" mapstructure:"accounts"`
	}
	Artifacts struct {
		Github struct {
			Enabled  bool
			Accounts []GithubArtifactAccount
		}
	}
}

type GithubArtifactAccount struct {
	Name     string
	Username string
	Token    string
}

// DockerAccount settings
type DockerAccount struct {
	Name         string `json:"name" mapstructure:"name"`
	Username     string `json:"username" mapstructure:"username"`
	PasswordFile string `json:"passwordFile" mapstructure:"passwordFile"`
	Address      string
	Repositories []string
}

// KubernetesAccount settings
type KubernetesAccount struct {
	Name       string   `json:"name" mapstructure:"name"`
	Namespaces []string `json:"namespaces" mapstructure:"namespaces"`
	// kubeconfig does not appear in clouddriver.yaml -- but should be stored as a secret
	KubeonfigFile    string `json:"kubeconfigFile" mapstructure:"kubeconfigFile"`
	DockerRegistries []DockerRegistry
	ProviderVersion  string
}

type DockerRegistry struct {
	AccountName string
}

type AWSAccount struct {
	Name      string
	AccountID string
	Regions   []AWSRegion
}

type AWSRegion struct {
	Name string
}
