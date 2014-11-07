package dockerwrapper_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDockerwrapper(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Dockerwrapper Suite")
}
