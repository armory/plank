/*
 * Copyright 2019 Armory, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License")
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package plank

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
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

	c := New(WithClient(client))
	assert.NotNil(t, c)
	val := map[string]string{}
	err := c.Get("/", &val)
	assert.Nil(t, err)
	assert.Equal(t, "value1", val["key1"])
}

func TestDefaultClient(t *testing.T) {
	client := New()
	assert.NotNil(t, client)
}

func TestURLMapCopy(t *testing.T) {
	client := New()
	assert.NotNil(t, client)
	client.URLs["orca"] = "foobar"
	assert.NotEqual(t, DefaultURLs["orca"], "foobar")
}

func TestOptions(t *testing.T) {
	test_transport := New(WithTransport(&http.Transport{MaxIdleConns:5}))
	assert.Equal(t, &http.Transport{MaxIdleConns: 5}, test_transport.http.Transport)

	test_client := New(WithClient(&http.Client{}))
	assert.Equal(t, &http.Client{}, test_client.http)

	test_retry_inc := New(WithRetryIncrement(5 * time.Second))
	assert.Equal(t, 5*time.Second, test_retry_inc.retryIncrement)

	test_fiat := New(WithFiatUser("foo"))
	assert.Equal(t, "foo", test_fiat.FiatUser)

	test_max_retries := New(WithMaxRetries(5))
	assert.Equal(t, 5, test_max_retries.maxRetry)

	test_with_urls := New(WithURLs(map[string]string{"foo":"http://foo"}))
	assert.Equal(t, map[string]string{"foo":"http://foo"}, test_with_urls.URLs)
}