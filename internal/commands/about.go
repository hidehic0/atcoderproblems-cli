package commands

import (
	"time"

	"github.com/spf13/cobra"
)

var AboutCmd = &cobra.Command{
	Use:   "about",
	Short: "Get user rated point sum and AC count",
	Long:  "Get user rated point sum and AC count",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		username := args[0]

		// AC Count
		AcCountCmd.SetArgs([]string{username, "-r"})
		if err := AcCountCmd.Execute(); err != nil {
			return err
		}

		time.Sleep(1 * time.Second) // API rate limit
		// Rated Point Sum
		RatedPointSumCmd.SetArgs([]string{username, "-r"})
		if err := RatedPointSumCmd.Execute(); err != nil {
			return err
		}

		return nil
	},
}
