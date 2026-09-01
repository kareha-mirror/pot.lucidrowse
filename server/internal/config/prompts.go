package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Prompts struct {
	Common            string `yaml:"common"`
	Areas             string `yaml:"areas"`
	NewFlavor         string `yaml:"new-flavor"`
	NewFlavorImage    string `yaml:"new-flavor-image"`
	UpdateFlavor      string `yaml:"update-flavor"`
	UpdateFlavorImage string `yaml:"update-flavor-image"`
	Events            string `yaml:"events"`
	NewAction         string `yaml:"new-action"`
	NewActionImage    string `yaml:"new-action-image"`
	UpdateAreaState   string `yaml:"update-area-state"`
	UpdateRegionState string `yaml:"update-region-state"`
	UpdateWorldState  string `yaml:"update-world-state"`
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
