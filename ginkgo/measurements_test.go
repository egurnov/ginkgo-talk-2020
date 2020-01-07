package ginkgo_test

import (
	"math/rand"
	"sync"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Measurements", func() {
	const N = 100

	Context("goroutines", func() {
		var queue chan struct{}
		var worker = func(count *int, queue chan struct{}, wg *sync.WaitGroup) {
			for range queue {
				*count++
				time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
			}
			wg.Done()
		}

		BeforeEach(func() {
			queue = make(chan struct{}, N)
			for i := 0; i < N; i++ {
				queue <- struct{}{}
			}
			close(queue)
		})

		Measure("Message throughput with one goroutine", func(b Benchmarker) {
			var countA int

			runtime := b.Time("runtime", func() {
				wg := sync.WaitGroup{}
				wg.Add(1)
				go worker(&countA, queue, &wg)
				wg.Wait()
			})

			Expect(runtime.Seconds()).Should(BeNumerically("<", 3.0))

			b.RecordValue("countA", float64(countA))
		}, 5)

		Measure("Message throughput with two goroutines", func(b Benchmarker) {
			var countA int
			var countB int

			runtime := b.Time("runtime", func() {
				wg := sync.WaitGroup{}
				wg.Add(2)
				go worker(&countA, queue, &wg)
				go worker(&countB, queue, &wg)
				wg.Wait()
			})

			Expect(runtime.Seconds()).Should(BeNumerically("<", 3.0))

			b.RecordValue("countA", float64(countA))
			b.RecordValue("countB", float64(countB))
		}, 5)
	})
})
