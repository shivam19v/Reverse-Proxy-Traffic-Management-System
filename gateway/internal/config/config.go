package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Config represents gateway configuration
type Config struct {
	Routes []Route `yaml:"routes"`
}

// Route represents a route configuration
type Route struct {
	Path    string   `yaml:"path"`
	Targets []string `yaml:"targets"`
}

// LoadConfig loads YAML config
func LoadConfig(path string) (*Config, error) {

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}


// package config

// import (
// 	"os"
// 	"errors"
// 	"gopkg.in/yaml.v3"
// )

// // Route represents one route in YAML
// type Route struct {
// 	Path   string `yaml:"path"`
// 	Target string `yaml:"target"`
// }

// // Config represents full YAML structure
// type Config struct {
// 	Routes []Route `yaml:"routes"`
// }

// // LoadConfig reads YAML config file
// func LoadConfig(path string) (*Config, error) {

// 	// Read YAML file
// 	data, err := os.ReadFile(path)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var config Config

// 	// Convert YAML → Go struct
// 	err = yaml.Unmarshal(data, &config)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if len(config.Routes) == 0 {
// 		return nil, errors.New("no routes found")
// 	}

// 	for _, route := range config.Routes {

// 		if route.Path == "" {
// 			return nil, errors.New("route path missing")
// 		}

// 		if route.Target == "" {
// 			return nil, errors.New("route target missing")
// 		}
// 	}

// 	return &config, nil
// }