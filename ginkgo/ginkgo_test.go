package ginkgo_test

import (
	"fmt"
	"log"
	"strconv"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Ginkgo", func() {

	Describe("test", func() {
		Context("inner", func() {
			When("I do this", func() {
				It("hurts", func() {
					fmt.Println("Ouch!")
				})
				Specify("test case", func() {
					fmt.Println("It's fine!")
				})
			})
		})
	})

	XIt("can fail unconditionally", func() {
		Fail("Noooooooo!")
	})

	writeCode := func() error { return nil }
	writeTests := func() error { return nil }
	qqq := "???"
	profit := ""

	Specify("my business plan", func() {
		By("Step 1")
		Expect(writeCode()).To(Succeed())

		By("Step 2")
		Expect(writeTests()).To(Succeed())

		By("Step 3")
		Expect(qqq).To(Equal("???"))

		By("Step 4")
		Expect(profit).ToNot(BeNil())
	})

	Context("Setup and Teardown", func() { // same as Describe
		var s string
		BeforeEach(func() { // creation
			log.Println("BeforeEach 1")

			s = "original"
		})

		JustBeforeEach(func() { // configuration
			log.Println("JustBeforeEach 1")

			log.Printf("Using s: %s\n", s)
		})

		JustAfterEach(func() { // diagnostic
			log.Println("JustAfterEach 1")
		})

		AfterEach(func() { // teardown
			log.Println("AfterEach 1")
		})

		Specify("outer context test", func() {
			log.Println("Test 1")

			log.Printf("s in test: %s\n", s)
		})

		Context("Inner context", func() {
			BeforeEach(func() { // can override creation, but reuse configuration
				log.Println("BeforeEach 2")

				s = "override"
			})

			JustBeforeEach(func() {
				log.Println("JustBeforeEach 2")

				log.Printf("Using s: %s\n", s)
			})

			JustAfterEach(func() {
				log.Println("JustAfterEach 2")
			})

			AfterEach(func() {
				log.Println("AfterEach 2")
			})

			It("works", func() {
				log.Println("Test 2")

				log.Printf("s in test: %s\n", s)
			})
		})
	})

	Context("Asynchronous functions", func() {
		XContext("in a bad case", func() {
			It("fails in a goroutine", func() {
				go func() {
					defer GinkgoRecover()
					time.Sleep(1 * time.Second)
					Fail("Oh noes!")
				}()
			})

			It("doesn't do anything bad", func() {
				time.Sleep(3 * time.Second)
			})
		})

		XWhen("done properly", func() {
			It("fails in a goroutine", func(done Done) {
				go func() {
					defer GinkgoRecover()
					time.Sleep(1 * time.Second)
					Fail("Oh noes!")
					close(done)
				}()
			}, 5) // timeout in seconds

			It("doesn't do anything bad", func() {
				time.Sleep(3 * time.Second)
			})
		})
	})

	Context("generating tests", func() {
		for i := 0; i < 10; i++ {
			i := i // create a local copy of the loop variable
			Specify("test #"+strconv.Itoa(i), func() {
				log.Println("Running test #", i)
			})
		}
	})
})
