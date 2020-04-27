package plate_test

import (
	"bytes"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/metalblueberry/plate"
)

var _ = Describe("Main", func() {

	testTemplate := func(file string) {
		templateFolder := "test_templates"
		input, err := os.Open(path.Join(templateFolder, file+".json"))
		Expect(err).To(BeNil())
		output := &bytes.Buffer{}

		cfg := plate.NewConfig()
		cfg.Input = input
		cfg.Output = output
		cfg.TemplateGlob = path.Join(templateFolder, "*.tmpl")
		cfg.TemplateToExecute = file + ".tmpl"

		err = plate.NewPlate(cfg).Run()
		Expect(err).To(BeNil())

		bytes, err := ioutil.ReadFile(path.Join(templateFolder, file+".txt"))
		Expect(err).To(BeNil())
		Expect(output.String()).To(Equal(string(bytes)))
	}

	files, err := filepath.Glob("test_templates/*.tmpl")
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		It("Should parse the template"+file, func() {
			fileName := strings.TrimPrefix(file, "test_templates/")
			fileWithoutExtension := strings.TrimSuffix(fileName, filepath.Ext(fileName))
			testTemplate(fileWithoutExtension)
		})
	}
})
