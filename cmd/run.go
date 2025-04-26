package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/RustedSilver/orcana/internal/api"
	"github.com/spf13/cobra"
)

var composeConfig []byte

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run compose orchestration",
	Long:  "Run compose orchestration with additional functionality",
	PreRun: func(cmd *cobra.Command, args []string) {
		//validation is performed on RootParent
		composeConfig, _ = exec.Command("docker", "compose", "config").Output()
	},
	Run: func(cmd *cobra.Command, args []string) {
		services, err := api.GetServices(composeConfig)
		if err != nil {
			fmt.Println("Could not get Services from compose")
			os.Exit(1)
		}

		inactiveServices := []string{"compose", "down"}
		activeServices := []string{"compose", "up", "-d"}

		for _, s := range services {
			if s.IsEnabled() {
				activeServices = append(activeServices, s.Name)
			} else {
				inactiveServices = append(inactiveServices, s.Name)
			}
		}

		if err := runWithArgs(activeServices); err != nil {
			fmt.Println("Could not start and run active services", err)
			os.Exit(1)
		}

		if len(inactiveServices) > 2 {
			if err := runWithArgs(inactiveServices); err != nil {
				fmt.Println("Could not disabled indicated services", err)
				os.Exit(1)
			}
		}

	},
}

func runWithArgs(s []string) error {
	c := exec.Command("docker", s...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	return c.Run()

}

func init() {
	rootCmd.AddCommand(runCmd)
}
