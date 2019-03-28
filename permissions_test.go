package plank

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetUser(t *testing.T) {
	payload := `{"name":"testapp","admin":true,"accounts":[],"applications":[]}}`
	client := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, req.URL.String(), "http://armory-fiat:7003/authorize/foo")
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(payload)),
			Header:     make(http.Header),
		}
	})

	c := New(client)
	assert.Nil(t, err)
	user, err := c.GetUser("foo")
	assert.Nil(t, err)
	assert.Equal(t, user.Name, "testapp")
	assert.True(t, user.IsAdmin())
	assert.Equal(t, len(user.Accounts), 0)
	assert.Equal(t, len(user.Applications), 0)
}

func TestHasAppWriteAccess(t *testing.T) {
	u := User{
		Applications: []Authorization{
			Authorization{Name: "foo", Authorizations: []string{"write"}},
			Authorization{Name: "bar", Authorizations: []string{"read-only"}},
		},
	}

	assert.True(t, u.HasAppWriteAccess("foo"))
	assert.False(t, u.HasAppWriteAccess("bar"))
	assert.False(t, u.HasAppWriteAccess("nonexistent"))
}
