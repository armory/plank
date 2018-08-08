package applications

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	tests := map[string]struct {
		A int
		B int

		Expected bool
	}{
		"Test Name": {
			Expected: true,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.True(t, test.Expected == true, "failed message")

			app := Application{
				Name: "hello",
			}
			taskId, _ := Create(app, nil)

			assert.True(t, taskId == "1234", "failed message")
		})
	}
}
