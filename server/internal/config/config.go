package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App struct {
		Mode string `yaml:"mode"`
		Addr string `yaml:"addr"`
	} `yaml:"app"`

	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
		SSLMode  string `yaml:"sslmode"`
	} `yaml:"database"`

	Vacuum struct {
		CheckEvery     int   `yaml:"check-every"`
		Threshold      int64 `yaml:"threshold"`
		ImageThreshold int64 `yaml:"image-threshold"`
	} `yaml:"vacuum"`

	AI struct {
		Agent       string  `yaml:"agent"`
		Key         string  `yaml:"key"`
		Temperature float64 `yaml:"temperature"`
		TopP        float64 `yaml:"top_p"`
	} `yaml:"ai"`

	PromptsPath string `yaml:"prompts-path"`

	Prompts *Prompts
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	cfg.Prompts, err = loadPrompts(cfg.PromptsPath)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
