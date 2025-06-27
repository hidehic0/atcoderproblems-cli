package cmd

import (
	"github.com/spf13/cobra"
	"hidehic0/atcoderproblems-cli/internal/commands"
)

var rootCmd = &cobra.Command{
	Use:  "atcoderproblems-cli",
	Long: "A tool that uses the AtCoder problems API to use the AtCoder problems function in CLI",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return cmd.Help()
		}
		return nil
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(commands.AcCountCmd)
	rootCmd.AddCommand(commands.RatedPointSumCmd)
	rootCmd.AddCommand(commands.AboutCmd)
	rootCmd.AddCommand(commands.RecommendationCmd)
}
