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
	user, err := c.GetUser("foo", "")
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

func generateTestClient(t *testing.T, resp interface{}, respCode int, targetEndpoint string) *Client{
	tc := NewTestClient(func(r *http.Request) *http.Response {
		assert.Equal(t, r.URL.String(), targetEndpoint)
		//TODO(ethanfrogers): assert that the X-Spinnaker-User header matches the incoming user
		b, err := json.Marshal(resp)
		if err != nil {
			t.Fatalf("failed to marshal response input: %s", err.Error())
		}
		return &http.Response{
			StatusCode: respCode,
			Body: ioutil.NopCloser(bytes.NewReader(b)),
			Header: make(http.Header),
		}
	})
	return New(WithClient(tc))
}

func TestClient_UserRoles(t *testing.T) {
	cases := map[string]struct{
		c *Client
		expectedOutput []string
		expectedErr error
		username string
	}{
		"happy path": {
			username: "armory",
			c: generateTestClient(t, []FiatRole{{Name: "team-a"}, {Name: "team-b"}}, http.StatusOK, "http://localhost:7003/authorize/armory/roles"),
			expectedOutput: []string{"team-a", "team-b"},
			expectedErr: nil,
		},
	}

	for testName, c := range cases {
		t.Run(testName, func(t *testing.T) {
			out, err := c.c.UserRoles(c.username, "")
			assert.Equal(t, c.expectedOutput, out)
			assert.Nil(t, err)
		})
	}
}

type mockFiatClient struct {
	rolesReturn []string
	errReturn error
}

func (m mockFiatClient) UserRoles(username, traceparent string) ([]string, error) {
	return m.rolesReturn, m.errReturn
}

type mockPermissable struct {
	permissions []string
	ReadPermissable
}

func (m mockPermissable) GetPermissions() []string {
	return m.permissions
}

func TestFiatPermissionEvaluator_HasReadPermission(t *testing.T) {
	cases := map[string]struct{
		mockClient mockFiatClient
		permissable ReadPermissable
		expectedResult bool
		orMode bool
	}{
		"contains all permissions": {
			expectedResult: true,
			mockClient: mockFiatClient{
				rolesReturn: []string{"team-a", "team-b"},
				errReturn:   nil,
			},
			permissable: mockPermissable{
				permissions:     []string{"team-a", "team-b"},
			},
		},
		"user is missing roles": {
			expectedResult: false,
			mockClient: mockFiatClient{
				rolesReturn: []string{"team-b"},
				errReturn:   nil,
			},
			permissable: mockPermissable{
				permissions:     []string{"team-a", "team-b"},
			},
		},
		"user has different roles": {
			expectedResult: false,
			mockClient: mockFiatClient{
				rolesReturn: []string{"team-c"},
				errReturn:   nil,
			},
			permissable: mockPermissable{
				permissions:     []string{"team-a", "team-b"},
			},
		},
		"permissions has different roles": {
			expectedResult: false,
			mockClient: mockFiatClient{
				rolesReturn: []string{"team-b"},
				errReturn:   nil,
			},
			permissable: mockPermissable{
				permissions:     []string{"team-c", "team-d"},
			},
		},
		"or mode - contains at least 1 role in common": {
			expectedResult: true,
			orMode: true,
			mockClient: mockFiatClient{
				rolesReturn: []string{"team-b", "team-a"},
				errReturn:   nil,
			},
			permissable: mockPermissable{
				permissions:     []string{"team-c", "team-d", "team-b"},
			},
		},
		"or mode - no overlapping permissions": {
			expectedResult: false,
			orMode: true,
			mockClient: mockFiatClient{
				rolesReturn: []string{"team-b", "team-a"},
				errReturn:   nil,
			},
			permissable: mockPermissable{
				permissions:     []string{"team-c", "team-d", "team-e"},
			},
		},
	}

	for testName, c := range cases {
		t.Run(testName, func(t *testing.T) {
			cfactory := func(opts ...ClientOption) FiatClient {
				return c.mockClient
			}
			evaluator := &FiatPermissionEvaluator{orMode: c.orMode, clientFactory: cfactory}
			res, _ := evaluator.HasReadPermission("test", "", c.permissable)
			assert.Equal(t, c.expectedResult, res)
		})
	}


}
