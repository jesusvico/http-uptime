package config

import (
	"fmt"
	"os"

	"github.com/jesusvico/http-uptime/internal/endpoint"
	"gopkg.in/yaml.v2"
)

// Struct to parse the YAML configuration file
type ConfigYaml struct {
	Endpoints []struct {
		Name   string `yaml:"name"`
		Url    string `yaml:"url"`
		Method string `yaml:"method"`
	} `yaml:"endpoints"`
}

// Struct to store the configuration
type Config struct {
	Endpoints map[string]endpoint.Endpoint
}

func New(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var cYaml ConfigYaml

	if err := yaml.Unmarshal(data, &cYaml); err != nil {
		return nil, err
	}

	// Create the Config struct from the ConfigYaml struct
	var c Config
	c.Endpoints = make(map[string]endpoint.Endpoint)
	for _, e := range cYaml.Endpoints {
		ep, err := endpoint.New(e.Name, e.Url, e.Method)
		if err != nil {
			return nil, err
		}
		if _, exists := c.Endpoints[e.Name]; exists {
			return nil, fmt.Errorf("Endpoint %s exists more than one time", e.Name)
		}
		c.Endpoints[e.Name] = *ep
	}

	return &c, nil
}
