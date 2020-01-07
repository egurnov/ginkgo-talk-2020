package ginkgo_test

// ginkgo bootstrap - generate test setup
// ginkgo blur/unfocus - replace all F* setup with normal versions

import (
	"log"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGinkgoTalk(t *testing.T) {
	log.SetOutput(GinkgoWriter)

	RegisterFailHandler(Fail)
	RunSpecs(t, "GinkgoTalk Suite")
}

var _ = BeforeSuite(func() {
	log.Println("Before Suite")
})

var _ = AfterSuite(func() {
	log.Println("After Suite")
})
