package config

import (
	"fmt"
	"net/url"
	"os"

	"gopkg.in/yaml.v3"
)

func Load(filepath string) (*Gateway, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Gateway

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse yaml: %w", err)
	}

	for i := range cfg.Routes {
		parsedUrl, err := url.Parse(cfg.Routes[i].Target)
		if err != nil {
			return nil, fmt.Errorf("invalid target url in config: %w", err)
		}
		cfg.Routes[i].ParsedURL = parsedUrl
	}

	return &cfg, nil
}
