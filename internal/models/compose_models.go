package models

import "errors"

type ComposeContext struct {
	Name     string             `yaml:"name"`
	Services map[string]Service `yaml:"services"`
}

type Service struct {
	Name          string
	ContainerName string `yaml:"container_name"`
	Image         string `yaml:"image"`
	Enabled       string `yaml:"x-enabled"`
}

// A service is considered enabled unless specified otherwise.
// This way the behaviour of docker compose up and orcana should not differ if param is missing
func (s *Service) IsEnabled() bool {
	return s.Enabled == "true" || s.Enabled == ""
}

func (s *Service) Validate() error {
	if s.Enabled != "true" && s.Enabled != "false" && s.Enabled != "" {
		return errors.New("Invalid value. Param x-enabled can contain only true/false values")
	}
	return nil
}

func (s *Service) AddServiceName(name string) {
	s.Name = name
}
