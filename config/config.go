package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile   string            `yaml:"input_file"`
	OutputFile  string            `yaml:"output_file"`
	StartPage   int               `yaml:"start_page"`
	EndPage     int               `yaml:"end_page"`
	Headers     []string          `yaml:"headers"`
	StartColumn string            `yaml:"start_column"` 
	TextColumn  string            `yaml:"text_column"`  
	Patterns    map[string]string `yaml:"patterns"`
}


func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
