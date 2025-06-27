package commands

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"hidehic0/atcoderproblems-cli/internal/api"
)

var RatedPointSumCmd = &cobra.Command{
	Use:   "rps",
	Short: "Get rated point sum",
	Long:  "rps <username>",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		username := args[0]
		showRank, err := cmd.Flags().GetBool("showrank")
		if err != nil {
			return err
		}

		data := api.GetRatedPointSum(username)
		if data.Count == -1 {
			os.Exit(256)
			return nil
		}

		fmt.Print("Rated Point Sum: ")
		api.GreenString.Print(strconv.Itoa(data.Count))
		if showRank {
			fmt.Print(" Rank: ")
			api.GreenString.Print(strconv.Itoa(data.Rank) + "th")
		}
		fmt.Println()
		return nil
	},
}

func init() {
	RatedPointSumCmd.Flags().BoolP("showrank", "r", false, "show AC rank")
}
