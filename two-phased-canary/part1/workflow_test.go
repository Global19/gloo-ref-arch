package part1_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/gloo-ref-arch/two-phased-canary/part1"
	"github.com/solo-io/go-utils/testutils"
	"testing"
)

func TestTwoPhasedCanary(t *testing.T) {
	RegisterFailHandler(Fail)
	testutils.RegisterPreFailHandler(
		func() {
			testutils.PrintTrimmedStack()
		})
	testutils.RegisterCommonFailHandlers()
	RunSpecs(t, "Two Phased Canary Test Suite")
}

var _ = Describe("Two Phased Canary, Part 1", func() {
	testWorkflow := part1.GetTestWorkflow()

	BeforeSuite(func() {
		testWorkflow.Setup(".")
	})

	It("works", func() {
		testWorkflow.Run(".")
	})
})
