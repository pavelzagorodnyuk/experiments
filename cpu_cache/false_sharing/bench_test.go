package falsesharing_test

import (
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
)

// AdjacentCounters places both counters on the same 64-byte cache line,
// so concurrent updates from different cores bounce that line back and
// forth (false sharing).
type AdjacentCounters struct {
	a int64
	b int64
}

// PaddedCounters separates a and b onto different cache lines (128 bytes
// total) by surrounding each with 56 bytes of padding.
type PaddedCounters struct {
	a int64
	_ [7]int64
	b int64
	_ [7]int64
}

// runCounters runs b.N atomic increments on ptrA and ptrB concurrently
// from two goroutines pinned to two OS threads via GOMAXPROCS(2).
func runCounters(b *testing.B, ptrA, ptrB *int64) {
	prevProcs := runtime.GOMAXPROCS(2)
	defer runtime.GOMAXPROCS(prevProcs)

	var wg sync.WaitGroup
	wg.Add(2)
	b.ResetTimer()
	go func() {
		defer wg.Done()
		for i := 0; i < b.N; i++ {
			atomic.AddInt64(ptrA, 1)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < b.N; i++ {
			atomic.AddInt64(ptrB, 1)
		}
	}()
	wg.Wait()
}

func BenchmarkFalseSharing(b *testing.B) {
	counters := &AdjacentCounters{}
	runCounters(b, &counters.a, &counters.b)
}

func BenchmarkNormalOperation(b *testing.B) {
	counters := &PaddedCounters{}
	runCounters(b, &counters.a, &counters.b)
}
