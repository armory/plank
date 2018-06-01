package permissions

import (
	"github.com/armory/plank/client"
)

const fiatBaseURL = "http://fiat" // TODO: read from yamls
const fiatWritePerm = "WRITE"

var defaultClient = client.New() // TODO: package variable for this is just asking for trouble. There will most certainly be side-effects.

type user struct {
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

func getUser(name string) (user, error) {
	url := fiatBaseURL + "/authorize/" + name
	var u user
	err := defaultClient.Get(url, u)
	if err != nil {
		return u, err
	}
	return u, nil
}
