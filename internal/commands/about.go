package commands

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"hidehic0/atcoderproblems-cli/internal/api"
)

var AboutCmd = &cobra.Command{
	Use:   "about",
	Short: "Get user rated point sum and AC count",
	Long:  "Get user rated point sum and AC count",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		username := args[0]

		// AC Count
		acData := api.GetAcCount(username)
		if acData.Count == -1 {
			return fmt.Errorf("User not found")
		}
		fmt.Print("AC count: ")
		api.GreenString.Print(acData.Count)
		fmt.Print(" Rank: ")
		api.GreenString.Print(fmt.Sprintf("%dth", acData.Rank))
		fmt.Println()

		time.Sleep(1 * time.Second) // API rate limit

		// Rated Point Sum
		rpsData := api.GetRatedPointSum(username)
		if rpsData.Count == -1 {
			return fmt.Errorf("User not found")
		}
		fmt.Print("Rated Point Sum: ")
		api.GreenString.Print(rpsData.Count)
		fmt.Print(" Rank: ")
		api.GreenString.Print(fmt.Sprintf("%dth", rpsData.Rank))
		fmt.Println()

		return nil
	},
}
