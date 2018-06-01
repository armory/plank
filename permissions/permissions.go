// Package permissions helps determine permissions for users.
package permissions

// Admin returns whether or not a user is an admin.
func Admin(username string) (bool, error) {
	u, err := getUser(username)
	if err != nil {
		return false, err
	}
	return u.Admin == true, nil
}

// WriteAccess returns whether or not a user can write pipelines/configs/etc. for an app.
func WriteAccess(user, app string) (bool, error) {
	u, err := getUser(user)
	if err != nil {
		return false, err
	}
	for _, a := range u.Applications {
		if a.Name == app && contains(a.Authorizations, fiatWritePerm) {
			return true, nil
		}
	}
	return false, nil
}

func contains(a []string, s string) bool {
	for _, v := range a {
		if v == s {
			return true
		}
	}
	return false
}
