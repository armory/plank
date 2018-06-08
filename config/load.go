package config

import (
	"github.com/armory/go-yaml-tools/pkg/spring"
	"github.com/mitchellh/mapstructure"
)

// Load settings from yaml file.
func Load() (Settings, error) {
	var s Settings
	m, err := spring.LoadDefault([]string{"spinnaker"})
	if err != nil {
		return s, err
	}
	err = mapstructure.Decode(m, &s)
	if err != nil {
		return s, err
	}
	return s, nil
}
