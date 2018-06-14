package config

import (
	"github.com/armory/go-yaml-tools/pkg/spring"
	"github.com/mitchellh/mapstructure"
)

// Load settings from yaml file.
func Load(otherSettings ...string) (Settings, error) {
	var s Settings
	settingFiles := append([]string{"spinnaker"}, otherSettings...)
	m, err := spring.LoadDefault(settingFiles)
	if err != nil {
		return s, err
	}
	err = mapstructure.Decode(m, &s)
	if err != nil {
		return s, err
	}
	return s, nil
}
