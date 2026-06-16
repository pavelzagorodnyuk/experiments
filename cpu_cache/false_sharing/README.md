# False Sharing

**Theory:** the cache-coherency protocol invalidates an entire 64-byte
cache line whenever any core writes to it, even if two cores are
writing to *different* variables that merely happen to share a line.
This forces the line to bounce between cores' caches on every write —
"false sharing" — even though there's no real data dependency.

**Implementation:** `AdjacentCounters{a, b int64}` packs both counters
into the same 64-byte line. `PaddedCounters` surrounds each counter
with 56 bytes of padding (`_ [7]int64`) so `a` and `b` land on separate
lines. `runCounters` pins the process to 2 OS threads
(`runtime.GOMAXPROCS(2)`) and runs two goroutines, each doing `b.N`
`atomic.AddInt64` ops on its own counter, synchronized with a
`sync.WaitGroup`.

Note: `PaddedCounters`' line separation relies on the runtime's heap
allocation happening to align `a` to a cache-line boundary, which is
likely but not guaranteed by the language. Go 1.23+'s `//go:align 64`
struct tag would make this deterministic if the measured ratio looks
unstable across runs.

**Run:** `make bench`

**What to look for:** `BenchmarkFalseSharing` should be several times
slower (typically 3-10x) than `BenchmarkNormalOperation`.
