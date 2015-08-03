package randSequence_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestRandSequence(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RandSequence Suite")
}
