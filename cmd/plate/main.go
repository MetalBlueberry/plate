package main

import (
	"os"

	"github.com/metalblueberry/plate"
)

func main() {
	cfg := plate.NewConfig()
	cfg.Input = os.Stdin
	cfg.Output = os.Stdout

	err := plate.Run(cfg)
	if err != nil {
		panic(err)
	}
}
