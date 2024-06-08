package main

import (
	"os"
	"github.com/Amr-Shams/aoc2024/challenge/cmd"
)

func main() {
	if err := cmd.NewRootCommand().Execute(); err != nil {
		os.Exit(1)
	}
}
