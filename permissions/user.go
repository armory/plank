package permissions

// User is returned by Fiat's /authorize endpoint.
type User struct {
	Name         string          `json:"name"`
	Admin        bool            `json:"admin"`
	Accounts     []Authorization `json:"accounts"`
	Applications []Authorization `json:"applications"`
}

// Authorization describes permissinos for an account or application.
type Authorization struct {
	Name string `json:"name"`
	// Authorizations can be 'READ' 'WRITE'
	Authorizations []string `json:"authorizations"`
}

func getUser(g getter, name string) (User, error) {
	path := "/authorize/" + name
	var u User
	err := g.Get(path, &u)
	if err != nil {
		return u, err
	}
	return u, nil
}
