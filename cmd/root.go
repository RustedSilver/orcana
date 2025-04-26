package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "orcana",
	Short: "An orchestration tool on top of docker compose",
	Long:  "A tool for adding additional features to Docker Compose orchestration in single-node environments",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if _, err := exec.Command("docker", "compose").Output(); err != nil {
			return errors.New("Docker compose is not present in the system!\n")
		}

		if _, err := exec.Command("docker", "compose", "config").Output(); err != nil {
			return errors.New("Could not read compose config. Verify if a correct compose.[yaml|yml] file is present")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {

		fmt.Println(err.Error())
		os.Exit(1)
	}

}
