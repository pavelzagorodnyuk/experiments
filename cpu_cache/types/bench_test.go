package types_test

import (
	"math/rand"
	"testing"
	"unsafe"
)

const typesNumElements = 12_582_912

const typesSeed = 42

const (
	wordStructSize = int(unsafe.Sizeof(WordStruct{}))
	struct75Size   = int(unsafe.Sizeof(Struct75{}))
	struct50Size   = int(unsafe.Sizeof(Struct50{}))
)

// Element counts for random-access benchmarks: equal total byte size across all variants.
// Reference: typesNumElements × wordStructSize = 805,306,368 B.
const wordStructRandomN = typesNumElements
const struct75RandomN   = typesNumElements * wordStructSize / struct75Size
const struct50RandomN   = typesNumElements * wordStructSize / struct50Size

// WordStruct is the 100% baseline: 8 int64 fields, 64 bytes, 0 padding.
type WordStruct struct {
	A, B, C, D, E, F, G, H int64
}

// Struct75 is 75% of WordStruct's size: 4 int64 + 4 int32 = 48 bytes,
// already a multiple of 8 -> 0 padding.
type Struct75 struct {
	A, B, C, D int64
	E, F, G, H int32
}

// Struct50 is 50% of WordStruct's size: 8 int32 fields, 32 bytes, 0 padding.
type Struct50 struct {
	A, B, C, D, E, F, G, H int32
}

func TestTypesStructSizes(t *testing.T) {
	if got := unsafe.Sizeof(WordStruct{}); got != 64 {
		t.Errorf("sizeof(WordStruct) = %d, want 64", got)
	}
	if got := unsafe.Sizeof(Struct75{}); got != 48 {
		t.Errorf("sizeof(Struct75) = %d, want 48", got)
	}
	if got := unsafe.Sizeof(Struct50{}); got != 32 {
		t.Errorf("sizeof(Struct50) = %d, want 32", got)
	}
}

func buildWordStructSlice(n int) []WordStruct {
	slice := make([]WordStruct, n)
	for i := range slice {
		v := int64(i)
		slice[i] = WordStruct{A: v, B: v, C: v, D: v, E: v, F: v, G: v, H: v}
	}
	return slice
}

func buildStruct75Slice(n int) []Struct75 {
	slice := make([]Struct75, n)
	for i := range slice {
		slice[i] = Struct75{A: int64(i), B: int64(i), C: int64(i), D: int64(i), E: int32(i), F: int32(i), G: int32(i), H: int32(i)}
	}
	return slice
}

func buildStruct50Slice(n int) []Struct50 {
	slice := make([]Struct50, n)
	for i := range slice {
		slice[i] = Struct50{A: int32(i), B: int32(i), C: int32(i), D: int32(i), E: int32(i), F: int32(i), G: int32(i), H: int32(i)}
	}
	return slice
}

// buildTypesSequentialIndices returns [0, n) in order.
func buildTypesSequentialIndices(n int) []uint {
	indices := make([]uint, n)
	for i := range indices {
		indices[i] = uint(i)
	}
	return indices
}

// buildTypesShuffledIndices returns a permutation of [0, n) via Fisher-Yates.
func buildTypesShuffledIndices(n int) []uint {
	indices := buildTypesSequentialIndices(n)
	r := rand.New(rand.NewSource(typesSeed))
	for i := n - 1; i > 0; i-- {
		j := r.Intn(i + 1)
		indices[i], indices[j] = indices[j], indices[i]
	}
	return indices
}

func BenchmarkWordStructSequential(b *testing.B) {
	slice := buildWordStructSlice(typesNumElements)
	indices := buildTypesSequentialIndices(typesNumElements)
	b.ResetTimer()

	var sum int64
	for i := 0; i < b.N; i++ {
		for _, idx := range indices {
			e := slice[idx]
			sum += e.A + e.B + e.C + e.D + e.E + e.F + e.G + e.H
		}
	}
	typesSink = sum
}

func BenchmarkWordStructRandom(b *testing.B) {
	slice := buildWordStructSlice(wordStructRandomN)
	indices := buildTypesShuffledIndices(wordStructRandomN)
	b.ResetTimer()

	var sum int64
	for i := 0; i < b.N; i++ {
		for _, idx := range indices {
			e := slice[idx]
			sum += e.A + e.B + e.C + e.D + e.E + e.F + e.G + e.H
		}
	}
	typesSink = sum
}

func BenchmarkStruct75Sequential(b *testing.B) {
	slice := buildStruct75Slice(typesNumElements)
	indices := buildTypesSequentialIndices(typesNumElements)
	b.ResetTimer()

	var sum int64
	for i := 0; i < b.N; i++ {
		for _, idx := range indices {
			e := slice[idx]
			sum += e.A + e.B + e.C + e.D + int64(e.E) + int64(e.F) + int64(e.G) + int64(e.H)
		}
	}
	typesSink = sum
}

func BenchmarkStruct75Random(b *testing.B) {
	slice := buildStruct75Slice(struct75RandomN)
	indices := buildTypesShuffledIndices(struct75RandomN)
	b.ResetTimer()

	var sum int64
	for i := 0; i < b.N; i++ {
		for _, idx := range indices {
			e := slice[idx]
			sum += e.A + e.B + e.C + e.D + int64(e.E) + int64(e.F) + int64(e.G) + int64(e.H)
		}
	}
	typesSink = sum
}

func BenchmarkStruct50Sequential(b *testing.B) {
	slice := buildStruct50Slice(typesNumElements)
	indices := buildTypesSequentialIndices(typesNumElements)
	b.ResetTimer()

	var sum int64
	for i := 0; i < b.N; i++ {
		for _, idx := range indices {
			e := slice[idx]
			sum += int64(e.A) + int64(e.B) + int64(e.C) + int64(e.D) + int64(e.E) + int64(e.F) + int64(e.G) + int64(e.H)
		}
	}
	typesSink = sum
}

func BenchmarkStruct50Random(b *testing.B) {
	slice := buildStruct50Slice(struct50RandomN)
	indices := buildTypesShuffledIndices(struct50RandomN)
	b.ResetTimer()

	var sum int64
	for i := 0; i < b.N; i++ {
		for _, idx := range indices {
			e := slice[idx]
			sum += int64(e.A) + int64(e.B) + int64(e.C) + int64(e.D) + int64(e.E) + int64(e.F) + int64(e.G) + int64(e.H)
		}
	}
	typesSink = sum
}

// typesSink prevents the compiler from eliminating the summation loops.
var typesSink int64
