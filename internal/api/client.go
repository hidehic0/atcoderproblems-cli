package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/fatih/color"
)

var RedString = color.New(color.FgRed).Add(color.Bold)

func fetchAPI(url string, target interface{}) error {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(body, &target); err != nil {
		RedString.Println("Error unmarshaling response:", err)
		return err
	}
	return nil
}
