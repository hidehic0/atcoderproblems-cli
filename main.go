package main

import (
	"os"

	"hidehic0/atcoderproblems-cli/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
