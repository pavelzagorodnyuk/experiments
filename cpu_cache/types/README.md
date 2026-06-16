# Element Density (50% / 75% / 100% of Word Size)

**Theory:** mirrors Experiment 2 from the other direction — instead of
adding padding to a fixed-size struct, this shrinks the struct itself
(using narrower field types) so a fixed-size cache or cache line holds
proportionally more elements. `WordStruct` (all `int64`) is the 100%
baseline.

**Implementation:** all three structs declare 8 distinct fields and are
zero-padding by construction (descending field sizes summing to a
multiple of 8):

| Struct      | Fields                                          | Size | % of WordStruct |
|-------------|--------------------------------------------------|------|------------------|
| `WordStruct`| 8x `int64`                                        | 64B  | 100%             |
| `Struct75`  | 4x `int64`, 4x `int32`                            | 48B  | 75%              |
| `Struct50`  | 2x `int64`, 3x `int32`, 1x `int16`, 2x `int8`     | 32B  | 50%              |

Sizes are asserted exactly via `unsafe.Sizeof` in `TestTypesStructSizes`.
All three use the same element count (`16,777,216`) as Experiment 2.
Each struct gets both a sequential and a random-traversal (shared
shuffled `indices`, seed 42) benchmark.

**Run:** `make bench`

**What to look for:** the same linear-vs-threshold question as
Experiment 2, here driven by density rather than waste — does halving
the struct size roughly halve the cache-miss-driven overhead?
