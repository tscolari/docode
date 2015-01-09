package docode_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDocode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Docode Suite")
}
