package plank

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetPipelines(t *testing.T) {
	client := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, req.URL.String(), "http://armory-front50:8080/pipelines/myapp")
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString("[]")),
			Header:     make(http.Header),
		}
	})

	c := New(client)
	val, err := c.GetPipelines("myapp")
	assert.Nil(t, err)
	assert.Equal(t, len(val), 0) // Should get 0 pipelines back.
}
