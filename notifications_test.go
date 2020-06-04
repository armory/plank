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