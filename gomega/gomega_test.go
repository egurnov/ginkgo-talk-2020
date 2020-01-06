package gomega_test

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestGomegaStandalone(t *testing.T) {
	g := NewGomegaWithT(t)
	g.Expect(5).To(Equal(5))
}
