package access_test

import (
	"math/rand"
	"runtime"
	"testing"
)

// numElements is large enough that a slice of uint (8 bytes/element) is
// 8x the Ryzen 9700X's 32MB L3, forcing every traversal to miss the cache.
const accessNumElements = 33_554_432

const accessSeed = 42

// buildSequential returns a circular linked list expressed as next-indices:
// slice[i] = i+1, with the last element wrapping to 0.
func buildSequential(n int) []uint {
	slice := make([]uint, n)
	for i := 0; i < n-1; i++ {
		slice[i] = uint(i + 1)
	}
	slice[n-1] = 0
	return slice
}

// buildRandom returns a single random cycle over [0, n) via Fisher-Yates,
// so that following slice[i] still visits every element exactly once
// per full traversal, but in memory-scattered order.
func buildRandom(n int) []uint {
	order := make([]uint, n)
	for i := range order {
		order[i] = uint(i)
	}
	r := rand.New(rand.NewSource(accessSeed))
	for i := n - 1; i > 0; i-- {
		j := r.Intn(i + 1)
		order[i], order[j] = order[j], order[i]
	}

	slice := make([]uint, n)
	for i := 0; i < n-1; i++ {
		slice[order[i]] = order[i+1]
	}
	slice[order[n-1]] = order[0]
	return slice
}

func BenchmarkSequentialAccess(b *testing.B) {
	slice := buildSequential(accessNumElements)
	b.ResetTimer()

	var next uint
	for i := 0; i < b.N; i++ {
		for j := 0; j < accessNumElements; j++ {
			next = slice[next]
		}
	}
	runtime.KeepAlive(next)
}

func BenchmarkRandomAccess(b *testing.B) {
	slice := buildRandom(accessNumElements)
	b.ResetTimer()

	var next uint
	for i := 0; i < b.N; i++ {
		for j := 0; j < accessNumElements; j++ {
			next = slice[next]
		}
	}
	runtime.KeepAlive(next)
}
