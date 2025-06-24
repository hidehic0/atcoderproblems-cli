package commands

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"hidehic0/atcoderproblems-cli/internal/api"
)

var AcCountCmd = &cobra.Command{
	Use:   "count",
	Short: "Get AC count",
	Long:  "count <username>",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		username := args[0]
		showRank, err := cmd.Flags().GetBool("showrank")
		if err != nil {
			return err
		}

		data := api.GetAcCount(username)
		if data.Count == -1 {
			os.Exit(256)
			return nil
		}

		fmt.Print("AC count: ")
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
	AcCountCmd.Flags().BoolP("showrank", "r", false, "show AC rank")
}
