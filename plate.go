package plate

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

type Config struct {
	Input             io.Reader
	Output            io.Writer
	TemplateGlob      string
	TemplateToExecute string
}

// NewConfig creates a config with default values
func NewConfig() Config {
	return Config{
		TemplateGlob: "*.tmpl",
	}
}

type TemplateData struct {
	Params map[string]interface{}
}

func Run(conf Config) error {
	decoder := json.NewDecoder(conf.Input)
	inputData := make(map[string]interface{})
	decoder.Decode(&inputData)

	templateData := &TemplateData{
		Params: inputData,
	}

	tpl, err := template.New("empty").Funcs(sprig.TxtFuncMap()).ParseGlob(conf.TemplateGlob)
	if err != nil {
		return err
	}
	log.Print(tpl.DefinedTemplates())
	switch {
	case len(tpl.Templates()) == 1:
		for _, tpls := range tpl.Templates() {
			return tpls.Execute(conf.Output, templateData)
		}
	case len(tpl.Templates()) > 1 && conf.TemplateToExecute != "":
		return tpl.ExecuteTemplate(conf.Output, conf.TemplateToExecute, templateData)
	default:
		return errors.New("You must specify the template to render with parsing multiple template files.")
	}

	return nil
}
