package plank

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

func TestGet(t *testing.T) {
	mockBody := `
	{
		"key1": "value1",
		"key2": "value2"
	}`
	client := NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(mockBody)),
			Header:     make(http.Header),
		}
	})

	c := New(client)
	assert.Nil(t, err)
	val := map[string]string{}
	err = c.Get("/", &val)
	assert.Nil(t, err)
	assert.Equal(t, "value1", val["key1"])
}

func TestDefaultClient(t *testing.T) {
	client := New(nil)
	assert.Nil(t, err)
	assert.NotNil(t, client)
}

func TestURLMapCopy(t *testing.T) {
	client := New(nil)
	assert.Nil(t, err)
	client.URLs["orca"] = "foobar"
	assert.NotEqual(t, DefaultURLs["orca"], "foobar")
}
