package main

import (
	"os"

	"github.com/faddat/testchain/cmd/testchaind/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := cmd.Execute(rootCmd); err != nil {
		os.Exit(1)
	}
}
