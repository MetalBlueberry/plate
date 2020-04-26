package plate_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTmpl(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tmpl Suite")
}
