package main

import (
	"flag"
	"os"

	"github.com/metalblueberry/plate"
)

func main() {
	template := flag.String("template", "default", "template to render.")
	glob := flag.String("glob", "*.tmpl", "template loading glob pattern.")
	flag.Parse()

	cfg := plate.NewConfig()
	cfg.Input = os.Stdin
	cfg.Output = os.Stdout
	cfg.TemplateToExecute = *template
	cfg.TemplateGlob = *glob

	err := plate.NewPlate(cfg).Run()
	if err != nil {
		panic(err)
	}
}
