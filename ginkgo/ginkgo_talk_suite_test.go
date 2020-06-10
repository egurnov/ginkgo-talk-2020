package ginkgo_test

import (
	"log"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// ginkgo -skipMeasurements -r
// go test -count=1 ./...

func TestGinkgoTalk(t *testing.T) {
	// ginkgo -skipMeasurements -v ./ginkgo
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
