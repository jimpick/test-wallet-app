package main

import (
	"os"

	"github.com/jimpick/test-wallet-app/cmd"
)

func main() {
	defer func() {
		os.Exit(cmd.ExitCode)
	}()
	cmd.Execute()
}
