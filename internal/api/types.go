package api

type AcCountData struct {
	Count int `json:"count"`
	Rank  int `json:"rank"`
}

type RatedPointSumData struct {
	Count int `json:"count"`
	Rank  int `json:"rank"`
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
