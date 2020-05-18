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
		assert.Equal(t, req.URL.String(), "http://armory-front50:8080/pipelines/myapp")
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString("[]")),
			Header:     make(http.Header),
		}
	})

	c := New(WithClient(client))
	val, err := c.GetPipelines("myapp")
	assert.Nil(t, err)
	assert.Equal(t, len(val), 0) // Should get 0 pipelines back.
}

func TestPipeline_ValidateRefIds(t *testing.T) {
	tests := map[string]struct {
		stage  string
		expected error
	}{
		"refIds_happy_path" : {
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
									"2"
								],
								"type": "manualJudgment"
							}
						]
					}`, nil,
		},
		"refIds_mandatory" : {
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
					}`, errors.New("refId is a mandatory field for stages"),
		},
		"refIds_duplicated" : {
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
							},
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
					}`, errors.New("refId should be unique, currently two or more stages share the same refId"),
		},
		"requisiteStageRefIds_does_not_exists" : {
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
					}`, errors.New("requisiteStageRefIds: mj1 does not exists"),
		},
		"requisiteStageRefIds_with_same_refId" : {
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
					}`, errors.New("refId cannot be dependant of itself (requisiteStageRefIds)"),
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
			err = pipe.ValidateRefIds()
			assert.Equal(t, c.expected, err)
		})
	}
}