package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"hidehic0/atcoderproblems-cli/internal/api"
)

var RecommendationCmd = &cobra.Command{
	Use:   "recommendation",
	Short: "Get recommended problem",
	Long:  "Get recommended problem",
	RunE: func(cmd *cobra.Command, args []string) error {
		atcUsername, ok := os.LookupEnv("ATC_USERNAME")
		if !ok && len(args) == 0 {
			api.RedString.Println("Please set ATC_USERNAME or provide a username")
			os.Exit(256)
		}

		var username string
		if len(args) == 0 {
			username = atcUsername
		} else {
			username = args[0]
		}

		criterionPerformance := api.GetCriterionPerformance(username)
		fmt.Printf("Criterion performance is %d\n", criterionPerformance)
		return nil
	},
}
