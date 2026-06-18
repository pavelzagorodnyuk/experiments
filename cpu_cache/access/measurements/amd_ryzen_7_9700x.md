# Benchmark Results — AMD Ryzen 7 9700X

## System Specs

| Component | Details |
|-----------|---------|
| CPU | AMD Ryzen 7 9700X 8-Core Processor |
| RAM | 16 GB DDR5 |
| OS | Fedora 44 |
| Architecture | amd64 |

## Results

```
goos: linux
goarch: amd64
pkg: github.com/pavelzagorodnyuk/experiments/cpu_cache/access
cpu: AMD Ryzen 7 9700X 8-Core Processor             
BenchmarkSequentialAccess 	      39	  30022675 ns/op	       0 B/op	       0 allocs/op
BenchmarkSequentialAccess 	      39	  30029828 ns/op	       0 B/op	       0 allocs/op
BenchmarkSequentialAccess 	      39	  29995191 ns/op	       0 B/op	       0 allocs/op
BenchmarkSequentialAccess 	      39	  30007056 ns/op	       0 B/op	       0 allocs/op
BenchmarkSequentialAccess 	      39	  30006724 ns/op	       0 B/op	       0 allocs/op
BenchmarkRandomAccess     	       1	2381334779 ns/op	       0 B/op	       0 allocs/op
BenchmarkRandomAccess     	       1	2385968881 ns/op	       0 B/op	       0 allocs/op
BenchmarkRandomAccess     	       1	2364403167 ns/op	       0 B/op	       0 allocs/op
BenchmarkRandomAccess     	       1	2342996058 ns/op	       0 B/op	       0 allocs/op
BenchmarkRandomAccess     	       1	2359965793 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/pavelzagorodnyuk/experiments/cpu_cache/access	21.638s
```
