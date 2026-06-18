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
pkg: github.com/pavelzagorodnyuk/experiments/cpu_cache/false_sharing
cpu: AMD Ryzen 7 9700X 8-Core Processor             
BenchmarkFalseSharing-16       	100000000	        10.55 ns/op	       0 B/op	       0 allocs/op
BenchmarkFalseSharing-16       	100000000	        10.44 ns/op	       0 B/op	       0 allocs/op
BenchmarkFalseSharing-16       	100000000	        10.38 ns/op	       0 B/op	       0 allocs/op
BenchmarkFalseSharing-16       	100000000	        10.30 ns/op	       0 B/op	       0 allocs/op
BenchmarkFalseSharing-16       	100000000	        10.34 ns/op	       0 B/op	       0 allocs/op
BenchmarkNormalOperation-16    	335831406	         3.568 ns/op	       0 B/op	       0 allocs/op
BenchmarkNormalOperation-16    	336669301	         3.569 ns/op	       0 B/op	       0 allocs/op
BenchmarkNormalOperation-16    	336435654	         3.568 ns/op	       0 B/op	       0 allocs/op
BenchmarkNormalOperation-16    	336241245	         3.569 ns/op	       0 B/op	       0 allocs/op
BenchmarkNormalOperation-16    	335931559	         3.566 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/pavelzagorodnyuk/experiments/cpu_cache/false_sharing	13.067s
```
