package plank

import (
	"strings"
)

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

// Admin returns whether or not a user is an admin.
func (c *Client) IsAdmin(username string) (bool, error) {
	u, err := c.GetUser(username)
	if err != nil {
		return false, err
	}
	return u.Admin == true, nil
}

// WriteAccess returns whether or not a user can write pipelines/configs/etc. for an app.
func (c *Client) HasWriteAccess(username, app string) (bool, error) {
	u, err := c.GetUser(username)
	if err != nil {
		return false, err
	}
	for _, a := range u.Applications {
		if a.Name == app && containsLowerCase(a.Authorizations, "WRITE") {
			return true, nil
		}
	}
	return false, nil
}

func containsLowerCase(a []string, s string) bool {
	for _, v := range a {
		if strings.ToLower(v) == strings.ToLower(s) {
			return true
		}
	}
	return false
}

// GetUser gets a user by name.
func (c *Client) GetUser(name string) (*User, error) {
	var u User
	if err := c.Get(c.URLs["fiat"]+"/authorize/"+name, &u); err != nil {
		return nil, err
	}
	return &u, nil
}
