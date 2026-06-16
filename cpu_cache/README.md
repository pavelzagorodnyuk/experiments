# CPU Cache Experiments

Four Go micro-benchmarks that make CPU cache behavior visible and
measurable, targeting an **AMD Ryzen 7 9700X** (8 cores, 32MB L3 shared
cache, 64-byte cache lines).

Each experiment lives in its own subdirectory with its own `Makefile`
and `README.md`. Run `make bench` from any of them.

## Experiments

| Directory | Topic |
|-----------|-------|
| [access/](access/) | Random vs. sequential memory access (prefetcher) |
| [alignment/](alignment/) | Struct padding waste (0% / 25% / 50%) |
| [types/](types/) | Element density (50% / 75% / 100% of word size) |
| [false_sharing/](false_sharing/) | False sharing between cores |

## Benchmark design

Experiments that compare structs of different sizes use two separate invariants depending on access pattern:

- **Sequential benchmarks** use an equal *element count* across all variants. Every benchmark iterates the same number of times, so throughput differences reflect the cost of touching a larger struct, not a different amount of work.

- **Random-access benchmarks** use an equal *total byte size* across all variants. Cache hit rate in random access depends on how much of the working set fits in cache, which is a function of bytes, not element count. Fixing the byte footprint ensures all variants see the same cache pressure, so results isolate the per-access cost of a larger struct rather than mixing in a different miss rate.

The reference sizes and derived element counts are defined as compile-time constants in each package:

| Package | Reference size | Anchor struct |
|---------|---------------|---------------|
| `alignment` | 402,653,184 B (~384 MB) | `AlignedStruct` (24 B) |
| `types` | 805,306,368 B (~768 MB) | `WordStruct` (64 B) |

See `padded25RandomN`, `padded50RandomN` in `alignment/bench_test.go` and `wordStructRandomN`, `struct50RandomN` in `types/bench_test.go`.

## Prerequisites

- Go 1.22+ (developed against 1.26)
