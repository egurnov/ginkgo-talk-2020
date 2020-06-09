package ginkgo_test

import (
	"errors"
	"log"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
)

var _ = Describe("Gomega", func() {

	It("has different notations", func() {
		Expect(5).To(Equal(5))    // Deep equality, strict about types
		Expect(5).NotTo(Equal(3)) // Also with negations
		Expect(5).ToNot(Equal(3)) //
		Ω(5).Should(Equal(5))     // Another notation, Ω = Option + Z on MacOS
	})

	It("allows annotations", func() {
		Expect(5).To(Equal(5), "Basic math should work as expected")
		Expect(5).To(Equal(5), "Also formatted: Hello, %s!", "World")
		Expect(5).To(Equal(5), func() string { return "Even dynamic" })
	})

	It("can handle errors", func() {
		good := func() (string, error) {
			return "OK", nil
		}
		Expect(good()).To(Equal("OK")) // Gomega can check your errors for you

		bad := func() (string, error) {
			return "OK", errors.New("oops")
		}
		// Expect(bad()).To(Equal("OK")) // Non-nil values fail tests

		_, err := bad()
		Expect(err).To(HaveOccurred())

		winner := func() error { return nil }
		Expect(winner()).To(Succeed())

		err = errors.New("oops")
		Expect(err).To(MatchError(errors.New("oops"))) // reflect.DeepEqual()
		Expect(err).To(MatchError("oops"))             // Matches ACTUAL.Error()
		Expect(err).To(MatchError(ContainSubstring("oo")))
	})

	It("can handle panics", func() {
		Expect(func() {
			panic("Nooooooooo!")
		}).To(Panic())
	})

	It("has multiple ways to assert equivalence", func() {
		Expect(5).To(Equal(5)) // uses reflect.DeepEqual
		Expect(map[string]int{"a": 1, "b": 2}).
			To(Equal(map[string]int{"b": 2, "a": 1}))
		Expect([]int{1, 2, 3}).To(Equal([]int{1, 2, 3}))
		type Custom struct {
			i int
			d float64
		}
		Expect(Custom{1, 2.0}).To(Equal(Custom{1, 2.0}))

		type KindOfInt int
		const (
			Zero KindOfInt = iota
			One
			Two
		)

		Expect(Zero).To(BeEquivalentTo(0)) // Convert ACTUAL’s type to that of EXPECTED
		Expect(Zero).ToNot(Equal(0))
		Expect(5.1).To(BeEquivalentTo(5)) // Type casting gotcha
		Expect(5).ToNot(BeEquivalentTo(5.1))

		type AlmostTheSame struct {
			i int
			d float64
		}
		Expect(Custom{1, 2.0}).ToNot(Equal(AlmostTheSame{1, 2.0}))
		Expect(Custom{1, 2.0}).To(BeEquivalentTo(AlmostTheSame{1, 2.0}))

		p1 := &struct{ v int }{v: 5}
		p2 := p1
		Expect(p1).To(BeIdenticalTo(p2)) // point to the exact same location in memory
		Expect(p1).ToNot(BeIdenticalTo(&struct{ v int }{v: 5}))
		Expect(p1).ToNot(BeIdenticalTo(struct{ v int }{v: 5}))
	})

	It("is good at comparing numbers", func() {
		Expect(5).To(BeNumerically("<", 5.1))         // Supports ==, ~, >, >=, <, <=
		Expect(5).To(BeNumerically("~", 5.005, 1e-2)) // Comparison with a threshold
		d1 := time.Date(2020, time.February, 5, 19, 30, 0, 0, time.UTC)
		d2 := time.Date(2020, time.February, 5, 19, 34, 0, 0, time.UTC)
		Expect(d1).To(BeTemporally("~", d2, 5*time.Minute))
	})

	It("has type-specific matchers", func() {
		By("Pointers")
		var p *int
		Expect(p).To(BeNil())
		Expect("").To(BeZero())

		By("Booleans")
		Expect(true).To(BeTrue()) // Only for bool
		Expect(false).To(BeFalse())

		By("Strings")
		Expect("Golang").To(HavePrefix("Go"))
		Expect("Abracadabra").To(ContainSubstring("cad"))
		Expect("x-y=z").To(ContainSubstring("%v-%v", "x", "y"))
		Expect("address@example.com").To(MatchRegexp("[a-z]+@[a-z]+\\.[a-z]{2,}"))
		Expect("{\"a\": 1, \"b\": 2}").To(MatchJSON("{\"b\": 2, \"a\": 1}")) // Similar XML and YAML

		By("Channels")
		ch := make(chan int, 1)
		var v int
		ch <- 5
		Expect(ch).To(Receive(&v))
		Expect(ch).To(BeSent(7))
		Expect(ch).To(Receive(Equal(7)))
		close(ch)
		Expect(ch).To(BeClosed())
	})

	It("works with collections", func() {
		theSequence := []int{4, 8, 15, 16, 23, 42}
		Expect(theSequence).ToNot(BeEmpty())
		Expect(theSequence).To(HaveLen(6))
		Expect(theSequence).To(ContainElement(23))
		Expect(15).To(BeElementOf(theSequence))
		Expect(theSequence).To(ConsistOf(8, 16, 42, 23, 15, 4))
		Expect([]string{"abc", "def"}).To(
			ConsistOf(
				HavePrefix("ab"),
				HaveSuffix("ef"),
			),
		)

		shoppingList := map[string]int{"apples": 4, "tomatoes": 10, "milk": 1}
		Expect(shoppingList).To(HaveKey("apples"))
		Expect(shoppingList).To(HaveKeyWithValue("tomatoes", 10))
		Expect(shoppingList).To(ConsistOf(1, 4, 10)) // match against values
	})

	Context("can make async assertions", func() {
		startSlowProcess := func(d time.Duration) func() bool {
			start := time.Now()
			return func() bool {
				log.Println("Running... ", time.Since(start))
				return time.Now().After(start.Add(d))
			}
		}

		It("eventually", func() {
			Eventually(startSlowProcess(500 * time.Millisecond)).Should(BeTrue())
			// Eventually(startSlowProcess(3 * time.Second)).Should(BeTrue()) // Will fail
			Eventually(
				startSlowProcess(3*time.Second),
				5*time.Second,
				500*time.Millisecond,
			).Should(BeTrue())
		})

		It("consistently", func() {
			Consistently(startSlowProcess(3 * time.Second)).Should(BeFalse())
			// Consistently(
			// 	startSlowProcess(3*time.Second),
			// 	5*time.Second,
			// 	500*time.Millisecond,
			// ).Should(BeFalse()) // Will fail
		})

		It("works nice with channels", func() {
			ch := make(chan struct{})
			time.AfterFunc(800*time.Millisecond, func() {
				ch <- struct{}{}
			})

			Eventually(ch).Should(Receive())
		})
	})

	Context("Kung Fu zone", func() {

		It("can combine matchers", func() {
			Expect(5).To(
				And( // Optionally SatisfyAll()
					BeNumerically(">", 4),
					BeNumerically("<", 6),
				),
			)
			Expect(5).To(
				Or( // Optionally SatisfyAny()
					BeNumerically(">", 0),
					BeNumerically("<", 0),
				),
			)
			Expect(5).To(Not(BeNil())) // Can also negate a single matcher
		})

		type T struct {
			name, id string
		}
		getName := func(t T) string {
			return t.name
		}
		getID := func(t T) string {
			return t.id
		}

		It("can transform value under test", func() {
			Expect(T{"a", "1"}).To(
				WithTransform(
					func(t T) string {
						return t.name
					},
					Equal("a"),
				),
			)

			arr := []T{{"a", "1"}, {"b", "1"}, {"c", "3"}}
			Expect(arr).To(
				ContainElement(
					WithTransform(
						getName,
						Equal("a"),
					),
				),
			)
			Expect(arr).To(ContainElement(WithTransform(getID, Equal("3"))))
		})

		It("can create custom matchers", func() {
			withName := func(id string) types.GomegaMatcher {
				return WithTransform(getName, Equal("a"))
			}
			withID := func(id string) types.GomegaMatcher {
				return WithTransform(getID, Equal("3"))
			}

			arr := []T{{"a", "1"}, {"b", "1"}, {"c", "3"}}
			Expect(arr).To(ContainElement(withName("a")))
			Expect(arr).To(ContainElement(withID("3")))

			Expect(5).To(BeInRange(3, 6))
		})
	})
})

func BeInRange(a, b interface{}) types.GomegaMatcher {
	return And(
		BeNumerically(">", a),
		BeNumerically("<", b),
	)
}
