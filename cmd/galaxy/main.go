package main

import (
	"fmt"
	"os"

	"github.com/galaxy-team/spesbas-chain/cmd/galaxy/launcher"
)

func main() {
	if err := launcher.Launch(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
