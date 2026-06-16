# Struct Alignment / Padding Waste

**Theory:** Go lays out struct fields in declaration order and pads for
alignment; padding bytes still occupy cache-line and cache-capacity
budget without holding useful data. More padding means fewer real
elements fit per cache line and per fetched page, so a cache of a given
size holds proportionally less useful data.

**Implementation:** three structs share the same 8 named data fields
(`A`..`H`: `int64, int32, int32, int16, int16, int16, int8, int8`,
ordered descending so there's no *unintentional* padding) and differ
only in how much *explicit* trailing padding (`_ [N]int8`) is appended:

| Struct           | Data fields | Padding | Total size | Waste |
|------------------|-------------|---------|------------|-------|
| `AlignedStruct`   | 24 bytes    | 0       | 24 bytes   | 0%    |
| `PaddedStruct25`  | 24 bytes    | 8       | 32 bytes   | 25%   |
| `PaddedStruct50`  | 24 bytes    | 24      | 48 bytes   | 50%   |

Sizes are asserted exactly via `unsafe.Sizeof` in `TestAlignmentStructSizes`.
All three use the same element count (`16,777,216`) so only struct size —
not element count — varies between them. Each struct gets both a
sequential-traversal and a random-traversal (shared shuffled `indices`,
seed 42) benchmark, to see whether padding waste matters more under
cache-capacity pressure (random) or memory bandwidth pressure
(sequential).

**Run:** `make bench`

**What to look for:** ns/op increasing with waste %, and whether that
increase is roughly linear (25%->50% costs about the same as 0%->25%)
or has threshold behavior; compare the size of the effect between the
`Sequential` and `Random` variants.
