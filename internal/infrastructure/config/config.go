package config

import (
	"io/ioutil"
	configDomain "my-story-time-api/internal/domain/config"

	"github.com/goccy/go-yaml"
)

func NewConfig() *configDomain.Config {
	config := &configDomain.Config{}

	content, err := ioutil.ReadFile("config.yaml")

	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(content, config)

	if err != nil {
		panic(err)
	}

	return config
}
