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
	"errors"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetPipelines(t *testing.T) {
	client := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, req.URL.String(), "http://localhost:8080/pipelines/myapp")
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString("[]")),
			Header:     make(http.Header),
		}
	})

	c := New(WithClient(client))
	val, err := c.GetPipelines("myapp", "")
	assert.Nil(t, err)
	assert.Equal(t, len(val), 0) // Should get 0 pipelines back.
}

func TestGetPipelinesWithGate(t *testing.T) {
	client := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, req.URL.String(), "http://localhost:8084/plank/pipelines/myapp")
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString("[]")),
			Header:     make(http.Header),
		}
	})

	c := New(WithClient(client))
	c.UseGateEndpoints()
	val, err := c.GetPipelines("myapp", "")
	assert.Nil(t, err)
	assert.Equal(t, len(val), 0) // Should get 0 pipelines back.
}

func TestUpsertPipelines(t *testing.T) {
	client := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, req.Method, "POST")
		responseJSON := `{
		  "status": "SUCCEEDED",
          "ref": "some/ref/12345",
		  "variables": {
			"exception": null
		  },
		  "otherField": "otherValue"
		}`
		// Create an HTTP response based on the responseJSON string
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBufferString(responseJSON)),
			Header:     make(http.Header),
		}
	})

	c := New(WithClient(client))
	err := c.UpsertPipeline(Pipeline{}, "Test", "")
	assert.Nil(t, err)
}

func TestUpsertPipelinesNoId(t *testing.T) {
	client := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, req.Method, "POST")
		responseJSON := `{
		  "status": "SUCCEEDED",
          "ref": "some/ref/12345",
		  "variables": {
			"exception": null
		  },
		  "otherField": "otherValue"
		}`
		// Create an HTTP response based on the responseJSON string
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBufferString(responseJSON)),
			Header:     make(http.Header),
		}
	})

	c := New(WithClient(client))
	err := c.UpsertPipeline(Pipeline{}, "", "")
	assert.Nil(t, err)
}

func TestDeletePipelines(t *testing.T) {
	client := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, req.URL.String(), "http://localhost:8080/pipelines/test-App/appName")
		assert.Equal(t, req.Method, "DELETE")
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString("[]")),
			Header:     make(http.Header),
		}
	})

	c := New(WithClient(client))
	err := c.DeletePipeline(Pipeline{Application: "test-App", Name: "appName"}, "")
	assert.Nil(t, err)
}

func TestDeletePipelinesWithGate(t *testing.T) {
	client := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, req.URL.String(), "http://localhost:8084/plank/pipelines/test-App/appName")
		assert.Equal(t, req.Method, "DELETE")
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString("[]")),
			Header:     make(http.Header),
		}
	})

	c := New(WithClient(client))
	c.UseGateEndpoints()
	err := c.DeletePipeline(Pipeline{Application: "test-App", Name: "appName"}, "")
	assert.Nil(t, err)
}

func TestPipeline_ValidateRefIds(t *testing.T) {
	tests := map[string]struct {
		stage           string
		expectedError   []string
		expectedWarning []string
	}{
		"refIds_happy_path": {
			`{
						"stages": [
							{
								"failPipeline": true,
								"judgmentInputs": [],
								"name": "Manual Judgment 2",
								"notifications": [],
								"refId": "1",
								"requisiteStageRefIds": [
								],
								"type": "manualJudgment"
							},
							{
								"failPipeline": true,
								"judgmentInputs": [],
								"name": "Manual Judgment 2",
								"notifications": [],
								"refId": "2",
								"requisiteStageRefIds": [
									"1"
								],
								"type": "manualJudgment"
							},
							{
								"failPipeline": true,
								"judgmentInputs": [],
								"name": "Manual Judgment 2",
								"notifications": [],
								"refId": "mj2",
								"requisiteStageRefIds": [
									"2","1"
								],
								"type": "manualJudgment"
							}
						]
					}`,
			[]string{},
			[]string{},
		},
		"refIds_mandatory": {
			`{
						"stages": [
							{
								"failPipeline": true,
								"judgmentInputs": [],
								"name": "Manual Judgment 2",
								"notifications": [],
								"type": "manualJudgment"
							}
						]
					}`,
			[]string{},
			[]string{"RefId field not found in stage"},
		},
		"refIds_and_stageref_do_not_exists": {
			`{
						"stages": [
							{
								"failPipeline": true,
								"judgmentInputs": [],
								"name": "Manual Judgment 2",
								"notifications": [],
								"requisiteStageRefIds": [
									"mj1"
								],
								"type": "manualJudgment"
							}
						]
					}`,
			[]string{"Referenced stage mj1 cannot be found."},
			[]string{"RefId field not found in stage"},
		},
		"refIds_duplicated": {
			`{
						"stages": [
							{
								"failPipeline": true,
								"judgmentInputs": [],
								"name": "Manual Judgment 2",
								"notifications": [],
								"refId": "mj2",
								"type": "manualJudgment"
							},
							{
								"failPipeline": true,
								"judgmentInputs": [],
								"name": "Manual Judgment 2",
								"notifications": [],
								"refId": "mj2",
								"type": "manualJudgment"
							}
						]
					}`,
			[]string{"Duplicate stage refId mj2 field found"},
			[]string{},
		},
		"requisiteStageRefIds_does_not_exists": {
			`{
						"stages": [
							{
								"failPipeline": true,
								"judgmentInputs": [],
								"name": "Manual Judgment 2",
								"notifications": [],
								"refId": "mj2",
								"requisiteStageRefIds": [
									"mj1"
								],
								"type": "manualJudgment"
							}
						]
					}`,
			[]string{"Referenced stage mj1 cannot be found."},
			[]string{},
		},
		"requisiteStageRefIds_with_same_refId": {
			`{
						"stages": [
							{
								"failPipeline": true,
								"judgmentInputs": [],
								"name": "Manual Judgment 2",
								"notifications": [],
								"refId": "mj2",
								"requisiteStageRefIds": [
									"mj2"
								],
								"type": "manualJudgment"
							}
						]
					}`,
			[]string{"mj2 refers to itself. Circular references are not supported"},
			[]string{},
		},
		"warning_no_stages": {
			`{
						"stages": [
						]
					}`,
			[]string{},
			[]string{"Current pipeline has no stages"},
		},
	}

	for testName, c := range tests {
		t.Run(testName, func(t *testing.T) {
			pipe := &Pipeline{}
			// It looks horrible but it works
			d := make(map[string][]map[string]interface{})
			err := json.Unmarshal([]byte(c.stage), &d)
			if err != nil {
				assert.Equal(t, true, err)
			}
			pipe.Stages = d["stages"]
			result := pipe.ValidateRefIds()
			expectedValidation := ValidationResult{Errors: nil, Warnings: nil}
			for _, errorMessage := range c.expectedError {
				if errorMessage != "" {
					expectedValidation.Errors = append(expectedValidation.Errors, errors.New(errorMessage))
				}
			}
			for _, warningMessage := range c.expectedWarning {
				if warningMessage != "" {
					expectedValidation.Warnings = append(expectedValidation.Warnings, warningMessage)
				}
			}
			assert.Equal(t, expectedValidation, result)
		})
	}
}
