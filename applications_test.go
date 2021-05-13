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
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetApplication(t *testing.T) {
	payload := `{"name":"testapp","email":"foo@bar.com"}`
	client := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, req.URL.String(), "http://localhost:8080/v2/applications/foo")
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(payload)),
			Header:     make(http.Header),
		}
	})

	c := New(WithClient(client))
	app, err := c.GetApplication("foo", "")
	assert.Nil(t, err)
	assert.Equal(t, app.Name, "testapp")
	assert.Equal(t, app.Email, "foo@bar.com")
}

func TestGetApplicationWithGate(t *testing.T) {
	payload := `{"name":"testapp","email":"foo@bar.com"}`
	client := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, req.URL.String(), "http://localhost:8084/plank/v2/applications/foo")
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(payload)),
			Header:     make(http.Header),
		}
	})

	c := New(WithClient(client))
	c.UseGateEndpoints()
	app, err := c.GetApplication("foo", "")
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
		case "http://localhost:8083/ops":
			assert.Equal(t, req.Method, "POST")
			payload = postPayload
		case "http://localhost:8083/refstring":
			assert.Equal(t, req.Method, "GET")
			payload = pollTaskPayload
		case "http://localhost:8080/v2/applications/foo":
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

	c := New(WithClient(client))
	err := c.CreateApplication(&Application{Name: "foo", Email: "Bar"}, "")
	assert.Nil(t, err)
}


func TestApplicationMarshalJSON (t *testing.T){
	app := Application{
		Name:        "appname",
		Email:       "useremail",
		Description: "appdescription",
		AppMetadata: map[string]interface{}{"testkey" : "testval", "testkey2" : "testval2"},
	}
	jsonbytes, err := json.Marshal(app)
	if err != nil {
		t.Fail()
	}

	validate := make(map[string]interface{})
	json.Unmarshal(jsonbytes, &validate)

	if _, ok := validate["testkey"]; !ok{
		t.Fail()
	}
}