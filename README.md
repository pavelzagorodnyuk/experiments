# experiments

A collection of small, self-contained experiments — each one its own
directory with its own module/tooling.

- [cpu_cache/](cpu_cache/) — Go micro-benchmarks demonstrating CPU cache
  effects (access pattern, struct padding, element density, false
  sharing), runnable via `go test -bench` and `perf stat`.
