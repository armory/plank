package plank

import (
	"encoding/json"
	"fmt"
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

	var notification NotificationsType
	err := json.Unmarshal([]byte(payload), &notification)
	if err != nil {
		t.Fail()
	}
	notification.FillAppNotificationFields("test")
	fmt.Sprintf("%v", notification)
}