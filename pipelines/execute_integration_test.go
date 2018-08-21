// +build integration

package pipelines

import (
	"github.com/armory/plank/client"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExecute(t *testing.T) {
	c, _ := client.New()
	s := Service{
		client:  c,
		gateURL: "http://spinnaker.dev.armory.io:8084",
	}
	ref, err := s.Execute("armoryhellodeploy", "Wait")
	t.Log(ref)
	assert.Nil(t, err)
	assert.NotEmpty(t, ref.Ref)
}
