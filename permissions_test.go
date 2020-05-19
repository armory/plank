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
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	payload := `{"name":"testapp","admin":true,"accounts":[],"applications":[]}`
	client := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, req.URL.String(), "http://localhost:7003/authorize/foo")
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(payload)),
			Header:     make(http.Header),
		}
	})

	c := New(WithClient(client))
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
