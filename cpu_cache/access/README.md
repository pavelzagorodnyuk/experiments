# Random vs. Sequential Access

**Theory:** the CPU's hardware prefetcher predicts forward, sequential
memory access and pulls future cache lines in before they're needed.
Random access defeats this — every load is an unpredictable address, so
the prefetcher can't help and a large fraction of loads miss all the way
out to DRAM.

**Implementation:** a `[]uint` of `33,554,432` elements (256MB, 8x the
32MB L3) is treated as a circular linked list — `slice[i]` holds the
index of the "next" element. `buildSequential` lays this out in index
order (`0->1->2->...`); `buildRandom` lays the same cycle out across a
Fisher-Yates shuffle (seed 42), so both benchmarks visit every element
exactly once per traversal, just in a different physical order. Both
benchmarks pointer-chase (`next = slice[next]`) for `b.N` full
traversals; pointer-chasing (rather than a plain `for` loop) ensures
each load's address depends on the previous load's result, so the CPU
genuinely cannot prefetch ahead in the random case.

**Run:** `make bench`

**What to look for:** `BenchmarkRandomAccess` should be dramatically
slower (tens of times) than `BenchmarkSequentialAccess`.
