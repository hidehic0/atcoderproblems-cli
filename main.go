package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var RedString = color.New(color.FgRed).Add(color.Bold)
var GreenString = color.New(color.FgGreen).Add(color.Bold)

type AcCountData struct {
	Count int `json:"count"`
	Rank  int `json:"rank"`
}

func getAcCount(username string) AcCountData {
	url := "https://kenkoooo.com/atcoder/atcoder-api/v3/user/ac_rank?user=" + username
	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	var data AcCountData
	if err := json.Unmarshal(body, &data); err != nil {
		RedString.Println("User not found")
		data.Count = -1
	}

	return data
}

type RatedPointSumData struct {
	Count int `json:"count"`
	Rank  int `json:"rank"`
}

func getRatedPointSum(username string) RatedPointSumData {
	url := "https://kenkoooo.com/atcoder/atcoder-api/v3/user/rated_point_sum_rank?user=" + username
	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	var data RatedPointSumData
	if err := json.Unmarshal(body, &data); err != nil {
		RedString.Println("User not found")
		data.Count = -1
	}

	return data
}

var rootCmd = &cobra.Command{
	Use:  "atcoderproblems-cli",
	Long: "A tool that uses the atcoder problems API to use the atcoder problems function in CLI",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			cmd.Help()
		}

		return nil
	},
}

var acCountCmd = &cobra.Command{
	Use:   "count",
	Short: "Get ac count",
	Long:  "count <username>",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: アニメーションを追加する
		username := args[0]

		showrank, err := cmd.Flags().GetBool("showrank")
		if err != nil {
			return err
		}

		data := getAcCount(username)

		if data.Count == -1 {
			os.Exit(256)
			return nil
		}

		fmt.Print("AC count: ")
		GreenString.Print(strconv.Itoa(data.Count))

		if showrank {
			fmt.Print(" Rank: ")
			GreenString.Print(strconv.Itoa(data.Rank) + "th")
		}

		fmt.Println()

		return nil
	},
}

var RatedPointSumDataCmd = &cobra.Command{
	Use:   "rps",
	Short: "Get rated point sum",
	Long:  "rps <username>",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: アニメーションを追加する
		username := args[0]

		showrank, err := cmd.Flags().GetBool("showrank")
		if err != nil {
			return err
		}

		data := getRatedPointSum(username)

		if data.Count == -1 {
			os.Exit(256)
			return nil
		}

		fmt.Print("Ratad Point Sum: ")
		GreenString.Print(strconv.Itoa(data.Count))

		if showrank {
			fmt.Print(" Rank: ")
			GreenString.Print(strconv.Itoa(data.Rank) + "th")
		}

		fmt.Println()

		return nil
	},
}

var aboutCmd = &cobra.Command{
	Use:   "about",
	Short: "Get user ratedpointsum and ac count",
	Long:  "Get user ratedpointsum and ac count",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: アニメーションを追加する
		username := args[0]

		// AC Count
		rootCmd.SetArgs([]string{"count", username, "-r"})

		rootCmd.Execute()

		time.Sleep(1 * time.Second) // apiの規約により1秒待機

		// Rated Point Sum
		rootCmd.SetArgs([]string{"rps", username, "-r"})

		rootCmd.Execute()

		return nil
	},
}

type ProblemData map[string]struct {
	Slope            float64 `json:"slope"`
	Intercept        float64 `json:"intercept"`
	Variance         float64 `json:"variance"`
	Difficulty       int     `json:"difficulty"`
	Discrimination   float64 `json:"discrimination"`
	IrtLogLikelihood float64 `json:"irt_loglikelihood"`
	IrtUsers         int     `json:"irt_users"`
	IsExperimental   bool    `json:"is_experimental"`
}

func getProblems() ProblemData {
	url := "https://kenkoooo.com/atcoder/resources/problem-models.json"
	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	var data ProblemData
	if err := json.Unmarshal(body, &data); err != nil {
		RedString.Println(err)
	}

	return data
}

type UserHistoryData []struct {
	IsRated           bool   `json:"IsRated"`
	Place             int    `json:"Place"`
	OldRating         int    `json:"OldRating"`
	NewRating         int    `json:"NewRating"`
	Performance       int    `json:"Performance"`
	InnerPerformance  int    `json:"InnerPerformance"`
	ContestScreenName string `json:"ContestScreenName"`
	ContestName       string `json:"ContestName"`
	ContestNameEn     string `json:"ContestNameEn"`
	EndTime           string `json:"EndTime"`
}

func getUserHistory(username string) UserHistoryData {
	url := "https://atcoder.jp/users/" + username + "/history/json"
	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	var data UserHistoryData
	if err := json.Unmarshal(body, &data); err != nil {
		RedString.Println("User not found")
		os.Exit(256)
	}

	return data
}

func getCriterionPerformance(username string) int {
	userHistory := getUserHistory(username)

	if len(userHistory) == 0 {
		RedString.Println(username + " not found" + "or " + username + " has not participated in the contest")
		os.Exit(256)
	}

	var performances []int

	for _, v := range userHistory {
		performances = append(performances, v.Performance)
	}

	var criterion_performance int

	sort.Sort(sort.Reverse(sort.IntSlice(performances)))

	if len(performances) < 4 {
		criterion_performance = performances[len(performances)-1]
	} else {
		criterion_performance = performances[3]
	}

	return criterion_performance
}

var recommendationCmd = &cobra.Command{
	Use:   "recommendation",
	Short: "git recoomendation problem",
	Long:  "git recoomendation problem",
	RunE: func(cmd *cobra.Command, args []string) error {

		atcUsername, ok := os.LookupEnv("ATC_USERNAME")

		if !ok && len(args) == 0 {
			RedString.Println("Please set ATC_USERNAME or set username")
			os.Exit(256)
		}

		var username string

		if len(args) == 0 {
			username = atcUsername
		} else {
			username = args[0]
		}

		criterionPerformance := getCriterionPerformance(username)

		fmt.Printf("criterion performance is %d\n", criterionPerformance)

		return nil
	},
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(acCountCmd)
	rootCmd.AddCommand(RatedPointSumDataCmd)
	rootCmd.AddCommand(aboutCmd)
	rootCmd.AddCommand(recommendationCmd)

	acCountCmd.Flags().BoolP("showrank", "r", false, "show ac rank")
	RatedPointSumDataCmd.Flags().BoolP("showrank", "r", false, "show ac rank")
}
