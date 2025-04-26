package api

import (
	"os"
	"slices"
	"testing"

	"github.com/RustedSilver/orcana/internal/models"
	"github.com/stretchr/testify/assert"
)

const BASE_TEST_FILE = "../../tests/examples/compose.yaml"

func TestGetServices(t *testing.T) {
	content, err := os.ReadFile(BASE_TEST_FILE)
	if err != nil {
		t.Fatal(err)
	}

	services, err := GetServices(content)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 4, len(services))

	firstService := findFirstByName(services, "app_1")
	secondService := findFirstByName(services, "app_2")
	thirdService := findFirstByName(services, "app_3")
	fourthService := findFirstByName(services, "app_4")

	assert.Equal(t, "app_1", firstService.Name)
	assert.Equal(t, "orcana_1", firstService.ContainerName)
	assert.Equal(t, "registry/project/orcana:0.0.1", firstService.Image)
	assert.Equal(t, false, firstService.IsEnabled())

	assert.Equal(t, "app_2", secondService.Name)
	assert.Equal(t, "orcana_2", secondService.ContainerName)
	assert.Equal(t, "registry/project/orcana:0.0.2", secondService.Image)
	assert.Equal(t, true, secondService.IsEnabled())

	assert.Equal(t, "app_3", thirdService.Name)
	assert.Equal(t, "orcana_3", thirdService.ContainerName)
	assert.Equal(t, "registry/project/orcana:0.0.3", thirdService.Image)
	assert.Equal(t, true, thirdService.IsEnabled())

	assert.Equal(t, "app_4", fourthService.Name)
	assert.Equal(t, "orcana_4", fourthService.ContainerName)
	assert.Equal(t, "registry/project/orcana:0.0.4", fourthService.Image)
	assert.Equal(t, false, fourthService.IsEnabled())
}

func findFirstByName(services []models.Service, sname string) models.Service {
	found := slices.Collect(func(yield func(models.Service) bool) {
		for _, s := range services {
			if s.Name == sname && !yield(s) {
				return
			}
		}
	})

	if len(found) == 0 {
		return models.Service{}
	}

	return found[0]

}
