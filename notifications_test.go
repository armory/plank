/*
 * Copyright 2020 Armory, Inc.
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
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNotificationsType_FillAppNotificationFields(t *testing.T) {
	var payload = `{
            "slack": [
                {
                    "when": [
                        "pipeline.failed"
                    ],
                    "address": "jossuecito-dinghy-pipes-this-is-good"
                }
            ],
            "email": [
                {
                    "when": [
                        "pipeline.starting"
                    ],
                    "address": "test@test.commmmm",
                    "cc": "test@test2.commmmmm"
                }
            ]
        }`

	var expected = `{"application":"test","email":[{"address":"test@test.commmmm","cc":"test@test2.commmmmm","level":"application","type":"email","when":["pipeline.starting"]}],"slack":[{"address":"jossuecito-dinghy-pipes-this-is-good","level":"application","type":"slack","when":["pipeline.failed"]}]}`

	var notification NotificationsType
	err := json.Unmarshal([]byte(payload), &notification)
	if err != nil {
		t.Fail()
	}
	notification.FillAppNotificationFields("test")

	body, err := json.Marshal(notification)
	assert.Equal(t, string(body), expected)
}

func TestNotificationsType_ValidateAppNotification(t *testing.T) {
	var payload = `{
            "slack": [
                {
                    "when": [
                        "pipeline.failed"
                    ],
                    "address": "jossuecito-dinghy-pipes-this-is-good"
                }
            ],
            "email": [
                {
                    "when": [
                        "pipeline.starting"
                    ],
                    "address": "test@test.commmmm",
                    "cc": "test@test2.commmmmm"
                }
            ],
            "application" : "test"
        }`

	var notification NotificationsType
	err := json.Unmarshal([]byte(payload), &notification)
	if err != nil {
		t.Fail()
	}
	err = notification.ValidateAppNotification()
	if err != nil {
		t.Fail()
	}
}

func TestNotificationsType_ValidateAppNotification_Fail(t *testing.T) {
	var payload = `{
            "slack": [
                {
                    "when": [
                        "pipeline.failed"
                    ],
                    "address": "jossuecito-dinghy-pipes-this-is-good"
                }
            ],
            "email": 
                {
                    "when": [
                        "pipeline.starting"
                    ],
                    "address": "test@test.commmmm",
                    "cc": "test@test2.commmmmm"
                },
            "application" : "test"
        }`

	var notification NotificationsType
	err := json.Unmarshal([]byte(payload), &notification)
	if err != nil {
		t.Fail()
	}
	err = notification.ValidateAppNotification()
	if err == nil {
		t.Fail()
	}
}

func TestNotificationsType_ValidateAppNotification_nil(t *testing.T) {
	var notification NotificationsType
	err := notification.ValidateAppNotification()
	if err != nil {
		t.Fail()
	}
}