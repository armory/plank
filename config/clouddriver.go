package config

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
}

// DockerAccount settings
type DockerAccount struct {
	Name     string `json:"name" mapstructure:"name"`
	Username string `json:"username" mapstructure:"username"`
	// Password file is a boolean in clouddriver.yaml
	PasswordFile string `json:"passwordFile" mapstructure:"passwordFile"`
}

// KubernetesAccount settings
type KubernetesAccount struct {
	Name       string   `json:"name" mapstructure:"name"`
	Namespaces []string `json:"namespaces" mapstructure:"namespaces"`
	// kubeconfig does not appear in clouddriver.yaml -- but should be stored as a secret
	KubeConfig string `json:"kubeconfig" mapstructure:"kubeconfig"`
}
