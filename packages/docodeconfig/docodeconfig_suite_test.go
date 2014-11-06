package docodeconfig_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDocodeconfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Docodeconfig Suite")
}
