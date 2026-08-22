package config

import (
	"log"
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

	Filter struct {
		Agent       string  `yaml:"agent"`
		Key         string  `yaml:"key"`
		Temperature float64 `yaml:"temperature"`
		TopP        float64 `yaml:"top_p"`
	} `yaml:"filter"`

	PromptsPath string `yaml:"prompts-path"`

	Prompts *Prompts
}

func Load(path string) *Config {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		log.Fatalf("failed to parse config: %v", err)
	}

	cfg.Prompts = loadPrompts(cfg.PromptsPath)

	return &cfg
}
