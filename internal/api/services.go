package api

import (
	"github.com/RustedSilver/orcana/internal/models"
	"github.com/goccy/go-yaml"
)

func GetServices(composeYaml []byte) ([]models.Service, error) {
	var cs models.ComposeContext
	if err := yaml.Unmarshal(composeYaml, &cs); err != nil {
		return nil, err
	}

	var s []models.Service
	for name, service := range cs.Services {
		service.AddServiceName(name)
		s = append(s, service)
	}

	return s, nil

}
