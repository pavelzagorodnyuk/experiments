package alignment_test

import (
	"math/rand"
	"testing"
	"unsafe"
)

const alignmentNumElements = 16_777_216

const alignmentSeed = 42

const (
	alignedStructSize  = int(unsafe.Sizeof(AlignedStruct{}))
	padded25StructSize = int(unsafe.Sizeof(PaddedStruct25{}))
	padded50StructSize = int(unsafe.Sizeof(PaddedStruct50{}))
)

// Element counts for random-access benchmarks: equal total byte size across all variants.
// Reference: alignmentNumElements × alignedStructSize = 402,653,184 B.
const padded25RandomN = alignmentNumElements * alignedStructSize / padded25StructSize
const padded50RandomN = alignmentNumElements * alignedStructSize / padded50StructSize

// AlignedStruct packs its 8 fields in descending size order so that the
// total (8+4+4+2+2+2+1+1=24 bytes) is already a multiple of the struct's
// 8-byte alignment (driven by the int64 field): zero padding anywhere.
type AlignedStruct struct {
	A int64
	B int32
	C int32
	D int16
	E int16
	F int16
	G int8
	H int8
}

// PaddedStruct25 is AlignedStruct plus 8 bytes of explicit trailing
// padding: 8/32 = 25% waste exactly.
type PaddedStruct25 struct {
	A int64
	B int32
	C int32
	D int16
	E int16
	F int16
	G int8
	H int8
	_ [8]int8
}

// PaddedStruct50 is AlignedStruct plus 24 bytes of explicit trailing
// padding: 24/48 = 50% waste exactly.
type PaddedStruct50 struct {
	A int64
	B int32
	C int32
	D int16
	E int16
	F int16
	G int8
	H int8
	_ [24]int8
}

func TestAlignmentStructSizes(t *testing.T) {
	if got := unsafe.Sizeof(AlignedStruct{}); got != 24 {
		t.Errorf("sizeof(AlignedStruct) = %d, want 24", got)
	}
	if got := unsafe.Sizeof(PaddedStruct25{}); got != 32 {
		t.Errorf("sizeof(PaddedStruct25) = %d, want 32", got)
	}
	if got := unsafe.Sizeof(PaddedStruct50{}); got != 48 {
		t.Errorf("sizeof(PaddedStruct50) = %d, want 48", got)
	}
}

func buildAlignedSlice(n int) []AlignedStruct {
	slice := make([]AlignedStruct, n)
	for i := range slice {
		slice[i] = AlignedStruct{A: int64(i), B: int32(i), C: int32(i), D: int16(i), E: int16(i), F: int16(i), G: int8(i), H: int8(i)}
	}
	return slice
}

func buildPadded25Slice(n int) []PaddedStruct25 {
	slice := make([]PaddedStruct25, n)
	for i := range slice {
		slice[i] = PaddedStruct25{A: int64(i), B: int32(i), C: int32(i), D: int16(i), E: int16(i), F: int16(i), G: int8(i), H: int8(i)}
	}
	return slice
}

func buildPadded50Slice(n int) []PaddedStruct50 {
	slice := make([]PaddedStruct50, n)
	for i := range slice {
		slice[i] = PaddedStruct50{A: int64(i), B: int32(i), C: int32(i), D: int16(i), E: int16(i), F: int16(i), G: int8(i), H: int8(i)}
	}
	return slice
}

// buildSequentialIndices returns [0, n) in order.
func buildSequentialIndices(n int) []uint {
	indices := make([]uint, n)
	for i := range indices {
		indices[i] = uint(i)
	}
	return indices
}

// buildShuffledIndices returns a permutation of [0, n) via Fisher-Yates.
func buildShuffledIndices(n int) []uint {
	indices := buildSequentialIndices(n)
	r := rand.New(rand.NewSource(alignmentSeed))
	for i := n - 1; i > 0; i-- {
		j := r.Intn(i + 1)
		indices[i], indices[j] = indices[j], indices[i]
	}
	return indices
}

func BenchmarkAlignedStructSequential(b *testing.B) {
	slice := buildAlignedSlice(alignmentNumElements)
	indices := buildSequentialIndices(alignmentNumElements)
	b.ResetTimer()

	var sum int64
	for i := 0; i < b.N; i++ {
		for _, idx := range indices {
			e := slice[idx]
			sum += e.A + int64(e.B) + int64(e.C) + int64(e.D) + int64(e.E) + int64(e.F) + int64(e.G) + int64(e.H)
		}
	}
	alignmentSink = sum
}

func BenchmarkAlignedStructRandom(b *testing.B) {
	slice := buildAlignedSlice(alignmentNumElements)
	indices := buildShuffledIndices(alignmentNumElements)
	b.ResetTimer()

	var sum int64
	for i := 0; i < b.N; i++ {
		for _, idx := range indices {
			e := slice[idx]
			sum += e.A + int64(e.B) + int64(e.C) + int64(e.D) + int64(e.E) + int64(e.F) + int64(e.G) + int64(e.H)
		}
	}
	alignmentSink = sum
}

func BenchmarkPadded25StructSequential(b *testing.B) {
	slice := buildPadded25Slice(alignmentNumElements)
	indices := buildSequentialIndices(alignmentNumElements)
	b.ResetTimer()

	var sum int64
	for i := 0; i < b.N; i++ {
		for _, idx := range indices {
			e := slice[idx]
			sum += e.A + int64(e.B) + int64(e.C) + int64(e.D) + int64(e.E) + int64(e.F) + int64(e.G) + int64(e.H)
		}
	}
	alignmentSink = sum
}

func BenchmarkPadded25StructRandom(b *testing.B) {
	slice := buildPadded25Slice(padded25RandomN)
	indices := buildShuffledIndices(padded25RandomN)
	b.ResetTimer()

	var sum int64
	for i := 0; i < b.N; i++ {
		for _, idx := range indices {
			e := slice[idx]
			sum += e.A + int64(e.B) + int64(e.C) + int64(e.D) + int64(e.E) + int64(e.F) + int64(e.G) + int64(e.H)
		}
	}
	alignmentSink = sum
}

func BenchmarkPadded50StructSequential(b *testing.B) {
	slice := buildPadded50Slice(alignmentNumElements)
	indices := buildSequentialIndices(alignmentNumElements)
	b.ResetTimer()

	var sum int64
	for i := 0; i < b.N; i++ {
		for _, idx := range indices {
			e := slice[idx]
			sum += e.A + int64(e.B) + int64(e.C) + int64(e.D) + int64(e.E) + int64(e.F) + int64(e.G) + int64(e.H)
		}
	}
	alignmentSink = sum
}

func BenchmarkPadded50StructRandom(b *testing.B) {
	slice := buildPadded50Slice(padded50RandomN)
	indices := buildShuffledIndices(padded50RandomN)
	b.ResetTimer()

	var sum int64
	for i := 0; i < b.N; i++ {
		for _, idx := range indices {
			e := slice[idx]
			sum += e.A + int64(e.B) + int64(e.C) + int64(e.D) + int64(e.E) + int64(e.F) + int64(e.G) + int64(e.H)
		}
	}
	alignmentSink = sum
}

// alignmentSink prevents the compiler from eliminating the summation loops.
var alignmentSink int64
