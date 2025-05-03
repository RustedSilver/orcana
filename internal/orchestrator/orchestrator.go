package orchestrator

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/RustedSilver/orcana/internal/models"
	"github.com/RustedSilver/orcana/internal/utils"
)

type OrcanaOrchestrator struct {
	Config *models.OrcanaConfig
}

func NewOrcanaOrchestrator() (*OrcanaOrchestrator, error) {
	if _, err := exec.Command("docker", "compose").Output(); err != nil {
		return nil, errors.New("Docker compose is not present in the system!\n")
	}

	content, err := exec.Command("docker", "compose", "config").Output()
	if err != nil {
		return nil, errors.New("Could not read compose config. Verify if a correct compose.[yaml|yml] file is present")
	}

	cf, err := models.NewOrcanaOrchestrationConfig(content)
	if err != nil {
		return nil, err
	}

	return &OrcanaOrchestrator{
		Config: cf,
	}, nil

}

func (o *OrcanaOrchestrator) Run() {
	fmt.Printf("%s:\n\tTotal Services [%s]\n\t[%s] Enabled\n\t[%s] Disabled\n\n",
		utils.ColorText(utils.Magenta, "Orcana Run", utils.Bold),
		utils.ColorText(utils.Green, strconv.Itoa(len(o.Config.ServiceBank.Enabled)+len(o.Config.ServiceBank.Disabled))),
		utils.ColorText(utils.Green, strconv.Itoa(len(o.Config.ServiceBank.Enabled))),
		utils.ColorText(utils.Red, strconv.Itoa(len(o.Config.ServiceBank.Disabled))),
	)

	StartContainers(o.Config.GetEnabledServiceNames()...)

	if len(o.Config.ServiceBank.Disabled) > 0 {
		TearDownContainers(o.Config.GetDisabledServiceNames()...)
	}

	// runResult := make(chan bool)

}

// Start Containers. Removes orphans
func StartContainers(names ...string) error {
	s := append([]string{"compose", "up", "-d", "--remove-orphans"}, names...)
	return run(s...)
}

// Stops and removes running container in compose. It will removery every container in case no args are passed in method
func TearDownContainers(names ...string) error {
	s := append([]string{"compose", "down"}, names...)
	return run(s...)

	// return run(...)
}

func run(args ...string) error {
	c := exec.Command("docker", args...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	return c.Run()
}
