package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Prompts struct {
	Common          string `yaml:"common"`
	Areas           string `yaml:"areas"`
	Create          string `yaml:"create"`
	Image           string `yaml:"image"`
	Update          string `yaml:"update"`
	UpdateImage     string `yaml:"update-image"`
	Events          string `yaml:"events"`
	Action          string `yaml:"action"`
	ActionImage     string `yaml:"action-image"`
	UpdateAreaState string `yaml:"update-area-state"`
}

func loadPrompts(path string) (*Prompts, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var prompts Prompts
	if err := yaml.Unmarshal(data, &prompts); err != nil {
		return nil, err
	}

	return &prompts, nil
}
