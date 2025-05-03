package cmd

import (
	"github.com/RustedSilver/orcana/internal/orchestrator"
	"github.com/spf13/cobra"
)

var orca *orchestrator.OrcanaOrchestrator

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run compose orchestration",
	Long:  "Run compose orchestration with additional functionality",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		o, err := orchestrator.NewOrcanaOrchestrator()
		if err != nil {
			return err
		}

		orca = o
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		orca.Run()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
