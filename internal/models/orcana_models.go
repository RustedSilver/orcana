package models

import "github.com/goccy/go-yaml"

type OrcanaConfig struct {
	ServiceBank OrcanaServicesBank
}

type OrcanaServicesBank struct {
	Enabled  []Service
	Disabled []Service
}

func (o *OrcanaConfig) GetEnabledServices() []Service {
	return o.ServiceBank.Enabled
}

func (o *OrcanaConfig) GetDisabledServices() []Service {
	return o.ServiceBank.Disabled
}

func (o *OrcanaConfig) GetEnabledServiceNames() []string {
	var names []string

	for _, s := range o.ServiceBank.Enabled {
		names = append(names, s.Name)
	}

	return names
}

func (o *OrcanaConfig) GetDisabledServiceNames() []string {
	var names []string

	for _, s := range o.ServiceBank.Disabled {
		names = append(names, s.Name)
	}

	return names
}

func NewOrcanaOrchestrationConfig(composeYaml []byte) (*OrcanaConfig, error) {
	var cs ComposeContext
	if err := yaml.Unmarshal(composeYaml, &cs); err != nil {
		return nil, err
	}

	var enabledServices []Service
	var disabledServices []Service

	for name, service := range cs.Services {
		service.AddServiceName(name)
		if service.IsEnabled() {
			enabledServices = append(enabledServices, service)
		} else {
			disabledServices = append(disabledServices, service)
		}
	}

	osb := OrcanaServicesBank{
		Enabled:  enabledServices,
		Disabled: disabledServices,
	}

	return &OrcanaConfig{
		ServiceBank: osb,
	}, nil
}
