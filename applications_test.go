package plank

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetApplication(t *testing.T) {
	client := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, req.URL.String(), "http://armory-front50:8080/v2/applications/foo")
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString("")),
			Header:     make(http.Header),
		}
	})

	c, err := New(client)
	assert.Nil(t, err)
	val := Application{}
	err = c.GetApplication("foo", &val)
	assert.NotNil(t, err) // bad payload means an error was returned.
}
