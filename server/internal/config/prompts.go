package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Prompts struct {
	Filter   string `yaml:"filter"`
	Common   string `yaml:"common"`
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
