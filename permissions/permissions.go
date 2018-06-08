package permissions

import (
	"strings"
)

// Admin returns whether or not a user is an admin.
func (s *Service) Admin(username string) (bool, error) {
	u, err := getUser(s.client, username)
	if err != nil {
		return false, err
	}
	return u.Admin == true, nil
}

// WriteAccess returns whether or not a user can write pipelines/configs/etc. for an app.
func (s *Service) WriteAccess(user, app string) (bool, error) {
	u, err := getUser(s.client, user)
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
