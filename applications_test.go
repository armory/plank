package plank

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetApplication(t *testing.T) {
	payload := `{"name":"testapp","email":"foo@bar.com"}`
	client := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, req.URL.String(), "http://armory-front50:8080/v2/applications/foo")
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(payload)),
			Header:     make(http.Header),
		}
	})

	c, err := New(client)
	assert.Nil(t, err)
	app, err := c.GetApplication("foo")
	assert.Nil(t, err)
	assert.Equal(t, app.Name, "testapp")
	assert.Equal(t, app.Email, "foo@bar.com")
}

func TestCreateApp(t *testing.T) {
	postPayload := `{"ref":"/refstring"}`
	pollTaskPayload := `{"id":"foo","status":"sure","endTime":42}`
	appPayload := `{"name":"testapp","email":"foo@bar.com"}`
	client := NewTestClient(func(req *http.Request) *http.Response {
		var payload string
		switch req.URL.String() {
		case "http://armory-orca:8083/ops":
			assert.Equal(t, req.Method, "POST")
			payload = postPayload
		case "http://armory-orca:8083/refstring":
			assert.Equal(t, req.Method, "GET")
			payload = pollTaskPayload
		case "http://armory-front50:8080/v2/applications/foo":
			payload = appPayload
		default:
			assert.Fail(t, "Unexpected URL requested: "+req.URL.String())
		}
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(payload)),
			Header:     make(http.Header),
		}
	})

	c, err := New(client)
	assert.Nil(t, err)
	err = c.CreateApplication(&Application{Name: "foo", Email: "Bar"})
	assert.Nil(t, err)
}
