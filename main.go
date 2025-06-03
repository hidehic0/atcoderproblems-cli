package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

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
	Use:  "example",
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

		// Rated Point Sum
		rootCmd.SetArgs([]string{"rps", username, "-r"})

		rootCmd.Execute()

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

	acCountCmd.Flags().BoolP("showrank", "r", false, "show ac rank")
	RatedPointSumDataCmd.Flags().BoolP("showrank", "r", false, "show ac rank")
}
