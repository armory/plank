package kayenta

type Kayenta struct {
	Kayenta struct {
		Datadog struct {
			Enabled  bool
			Accounts []DataDogAccount
		}
	}
}

type DataDogAccount struct {
	Name            string
	APIKey          string
	ApplicationKey  string
	SupportTypes    []string
	EndPointBaseURL string `json:"endpoint.baseUrl"`
}
