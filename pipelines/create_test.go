package pipelines

import (
	"github.com/armory/plank/client"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreate(t *testing.T) {
	c, _ := client.New()
	s := Service{
		client:     c,
		front50URL: "http://spinnaker.dev.armory.io:8080",
	}
	p := Pipeline{
		Name:        "planktestpipeline",
		Application: "plankappcreationtest",
	}
	err := s.Create(p)
	assert.Nil(t, err)

}

func TestGet(t *testing.T) {
	c, _ := client.New()
	s := Service{
		client:     c,
		front50URL: "http://spinnaker.dev.armory.io:8080",
	}
	pipes, err := s.Get("armoryhellodeploy")
	t.Log(pipes)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, len(pipes))
}
