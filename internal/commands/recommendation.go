package commands

import (
	"fmt"
	"os"

	"github.com/fatih/color"
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

		problems := api.GetProblems()

		problem_quantity, err := cmd.Flags().GetInt("quantity")

		if err != nil {
			return err
		}

		is_test_tube, err := cmd.Flags().GetBool("Experimental")

		if err != nil {
			return nil
		}

		var left, right int = 0, 5000

		for right-left > 1 {
			mid := (left + right) / 2
			cnt := 0

			for _, problem := range problems {
				if criterionPerformance-mid <= problem.Difficulty && problem.Difficulty <= criterionPerformance+mid && (is_test_tube || !problem.IsExperimental) {
					cnt++
				}

			}

			if cnt < problem_quantity {
				left = mid
			} else {
				right = mid
			}
		}

		problemNameMap := api.GetProblemNameMap()

		boldPurpleString := color.New(color.Bold, color.FgHiMagenta)

		for problem_id, problem := range problems {
			if criterionPerformance-right <= problem.Difficulty && problem.Difficulty <= criterionPerformance+right && (is_test_tube || !problem.IsExperimental) {
				// fmt.Println(problemNameMap[problem_id], ":", api.GetProblemUrl(problem_id))
				fmt.Printf("Name: %s Difficulty: %d URL: %s\n", boldPurpleString.Sprint(problemNameMap[problem_id]), problem.Difficulty, boldPurpleString.Sprint(api.GetProblemUrl(problem_id)))
			}
		}

		return nil
	},
}

func init() {
	RecommendationCmd.Flags().IntP("quantity", "q", 10, "Number of questions to choose")
	RecommendationCmd.Flags().BoolP("Experimental", "e", false, "Include Experimental Difficulty problems")
}
