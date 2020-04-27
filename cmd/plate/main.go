package main

import (
	"os"

	"github.com/metalblueberry/plate"
)

func main() {
	cfg := plate.NewConfig()
	cfg.Input = os.Stdin
	cfg.Output = os.Stdout
	cfg.TemplateToExecute = "model_generator.tmpl"
	err := plate.NewPlate(cfg).Run()
	if err != nil {
		panic(err)
	}
}
