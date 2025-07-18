package api

import (
	"os"
	"sort"
	"strings"
	"time"

	"github.com/fatih/color"
)

var GreenString = color.New(color.FgGreen).Add(color.Bold)

func GetAcCount(username string) AcCountData {
	var data AcCountData
	url := "https://kenkoooo.com/atcoder/atcoder-api/v3/user/ac_rank?user=" + username
	if err := fetchAPI(url, &data); err != nil {
		RedString.Println("User not found")
		data.Count = -1
	}
	return data
}

func GetRatedPointSum(username string) RatedPointSumData {
	var data RatedPointSumData
	url := "https://kenkoooo.com/atcoder/atcoder-api/v3/user/rated_point_sum_rank?user=" + username
	if err := fetchAPI(url, &data); err != nil {
		RedString.Println("User not found")
		data.Count = -1
	}
	return data
}

func GetProblems() ProblemData {
	var data ProblemData
	url := "https://kenkoooo.com/atcoder/resources/problem-models.json"
	time.Sleep(time.Second * 1)
	if err := fetchAPI(url, &data); err != nil {
		RedString.Println("Failed to fetch problems")
	}
	return data
}

func GetProblemUrl(problemID string) string {
	contestName := strings.Split(problemID, "_")[0]
	return "https://atcoder.jp/contests/" + contestName + "/tasks/" + problemID
}

func GetProblemNameMap() map[string]string {
	type ProblemNameData []struct {
		ID    string `json:"id"`
		Title string `json:"title"`
	}

	url := "https://kenkoooo.com/atcoder/resources/problems.json"

	var data ProblemNameData

	time.Sleep(time.Second * 1)

	if err := fetchAPI(url, &data); err != nil {
		RedString.Println("Failed to fetch problems")
		os.Exit(256)
	}

	res := make(map[string]string)

	for _, v := range data {
		problemName := v.Title
		problemName = problemName[strings.Index(problemName, " ")+1:]
		res[v.ID] = problemName
	}

	return res
}

func GetUserHistory(username string) UserHistoryData {
	var data UserHistoryData
	url := "https://atcoder.jp/users/" + username + "/history/json"
	if err := fetchAPI(url, &data); err != nil {
		RedString.Println("User not found")
		os.Exit(256)
	}
	return data
}

func GetCriterionPerformance(username string) int {
	userHistory := GetUserHistory(username)
	if len(userHistory) == 0 {
		RedString.Println(username + " not found or has not participated in the contest")
		os.Exit(256)
	}

	var performances []int
	for _, v := range userHistory {
		performances = append(performances, v.Performance)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(performances)))
	if len(performances) < 4 {
		return performances[len(performances)-1]
	}
	return performances[3]
}
