package plate

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
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

type TemplateData map[string]interface{}

type Plate struct {
	Base *template.Template
	Conf Config
}

func NewPlate(conf Config) *Plate {
	return &Plate{
		Base: template.New("base"),
		Conf: conf,
	}
}

func (plate *Plate) Run() error {
	log.SetOutput(os.Stderr)
	decoder := json.NewDecoder(plate.Conf.Input)
	inputData := make(TemplateData)
	decoder.Decode(&inputData)

	tpl, err := plate.Base.
		Funcs(sprig.TxtFuncMap()).
		Funcs(
			template.FuncMap{
				"file":      plate.NewFile,
				"stemplate": plate.Stemplate,
				"parseJSON": plate.ParseJSON,
			}).
		ParseGlob(plate.Conf.TemplateGlob)

	if err != nil {
		return err
	}
	log.Print(tpl.DefinedTemplates())
	switch {
	case len(tpl.Templates()) == 1:
		for _, tpls := range tpl.Templates() {
			return tpls.Execute(plate.Conf.Output, inputData)
		}
	case len(tpl.Templates()) > 1 && plate.Conf.TemplateToExecute != "":
		return tpl.ExecuteTemplate(plate.Conf.Output, plate.Conf.TemplateToExecute, inputData)
	default:
		return errors.New("You must specify the template to render when having multiple templates defined.")
	}

	return nil
}

func (plate *Plate) NewFile(file string, template string, data interface{}) error {

	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	target, err := filepath.Abs(file)
	if err != nil {
		return err
	}

	if !strings.HasPrefix(target, dir) {
		return errors.New("Created files must be under cwd")
	}

	out, err := os.Create(target)
	if err != nil {
		return err
	}
	defer out.Close()
	return plate.Base.ExecuteTemplate(out, template, data)
}

func (plate *Plate) Stemplate(template string, data interface{}) (string, error) {
	buf := &bytes.Buffer{}
	err := plate.Base.ExecuteTemplate(buf, template, data)
	return buf.String(), err
}

func (plate *Plate) ParseJSON(template string, data interface{}) (interface{}, error) {
	buf := bytes.Buffer{}
	err := plate.Base.ExecuteTemplate(&buf, template, data)
	if err != nil {
		return nil, err
	}
	var out interface{}
	err = json.Unmarshal(buf.Bytes(), &out)
	if err != nil {
		log.Print(buf.String())
		return nil, fmt.Errorf("Error parsing string as json, cause %w", err)
	}
	return out, err
}
