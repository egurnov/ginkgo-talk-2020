package ginkgo_test

import (
	"math/rand"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = XDescribe("Flaky", func() {
	It("fails sometimes", func() {
		rand.Seed(time.Now().UnixNano())
		Expect(rand.Intn(100)).ToNot(BeNumerically("==", 42))
	})
})
