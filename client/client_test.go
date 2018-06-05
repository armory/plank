package client

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGet(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mockBody := `
		{
			"key1": "value1",
			"key2": "value2"
		}`
		fmt.Fprintln(w, mockBody)
	}))
	defer ts.Close()

	c, err := New(BaseURL(ts.URL))
	assert.Nil(t, err)
	val := map[string]string{}
	err = c.Get("/", &val)
	assert.Nil(t, err)
	assert.Equal(t, "value1", val["key1"])
}

func TestBaseURL(t *testing.T) {
	input := "http://fiat/"
	expected := "http://fiat"
	c, err := New(BaseURL(input))
	assert.Nil(t, err)
	assert.Equal(t, expected, c.baseURL)
}
