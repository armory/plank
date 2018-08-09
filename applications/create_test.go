package applications

import (
	"github.com/armory/plank/client"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	c, _ := client.New(client.MaxRetry(1))
	s := Service{
		orcaURL:    "http://spinnaker.dev.armory.io:8083",
		front50URL: "http://spinnaker.dev.armory.io:8080",
		client:     c,
		pollTime:   1 * time.Second,
	}
	name := "plankappcreationtest"
	email := "test@armory.io"
	a, err := s.Create(Application{Name: name, Email: email})
	assert.Nil(t, err)
	assert.Equal(t, name, a.Name)
	assert.Equal(t, email, a.Email)
}
