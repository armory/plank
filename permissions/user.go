package permissions

type User struct {
	Name         string          `json:"name"`
	Admin        bool            `json:"admin"`
	Accounts     []Authorization `json:"accounts"`
	Applications []Authorization `json:"applications"`
}

type Authorization struct {
	Name string `json:"name"`
	// Authorizations can be 'READ' 'WRITE'
	Authorizations []string `json:"authorizations"`
}

func getUser(name string) (User, error) {
	path := "/authorize/" + name
	var u User
	err := defaultFiatClient.Get(path, &u)
	if err != nil {
		return u, err
	}
	return u, nil
}
