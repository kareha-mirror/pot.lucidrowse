package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Prompts struct {
	Common      string `yaml:"common"`
	Create      string `yaml:"create"`
	Image       string `yaml:"image"`
	Update      string `yaml:"update"`
	UpdateImage string `yaml:"update-image"`
	Action      string `yaml:"action"`
	ActionImage string `yaml:"action-image"`
}

func loadPrompts(path string) *Prompts {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to load prompts: %v", err)
	}

	var prompts Prompts
	if err := yaml.Unmarshal(data, &prompts); err != nil {
		log.Fatalf("failed to parse prompts: %v", err)
	}

	return &prompts
}
