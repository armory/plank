// +build integration

package applications

import (
	"github.com/armory/plank/client"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"fmt"
)

func TestCreate(t *testing.T) {
	c, _ := client.New(client.MaxRetry(1))
	s := Service{
		orcaURL:    "http://spinnaker.dev.armory.io:8083",
		front50URL: "http://spinnaker.dev.armory.io:8080",
		client:     c,
		pollTime:   30 * time.Second, // dev is kind of slow
	}

	tf := "01021504" // MM DD HH mm
	name := fmt.Sprintf("plankappcreationtest%s", time.Now().Format(tf))
	email := "test@armory.io"
	a, err := s.Create(Application{Name: name, Email: email})
	assert.Nil(t, err)
	assert.Equal(t, name, a.Name)
	assert.Equal(t, email, a.Email)
}
