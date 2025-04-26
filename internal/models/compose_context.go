package models

type ComposeContext struct {
	Name     string             `yaml:"name"`
	Services map[string]Service `yaml:"services"`
}
