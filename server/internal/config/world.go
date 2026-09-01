package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type World struct {
	StateSeed string `yaml:"state-seed"`
}

func loadWorld(path string) (*World, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var world World
	if err := yaml.Unmarshal(data, &world); err != nil {
		return nil, err
	}

	return &world, nil
}
