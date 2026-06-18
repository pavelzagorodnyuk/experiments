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
pkg: github.com/pavelzagorodnyuk/experiments/cpu_cache/types
cpu: AMD Ryzen 7 9700X 8-Core Processor             
BenchmarkWordStructSequential 	      57	  19548201 ns/op	       0 B/op	       0 allocs/op
BenchmarkWordStructSequential 	      58	  19488317 ns/op	       0 B/op	       0 allocs/op
BenchmarkWordStructSequential 	      58	  19488932 ns/op	       0 B/op	       0 allocs/op
BenchmarkWordStructSequential 	      57	  19452067 ns/op	       0 B/op	       0 allocs/op
BenchmarkWordStructSequential 	      58	  19452405 ns/op	       0 B/op	       0 allocs/op
BenchmarkWordStructRandom     	       5	 214442882 ns/op	       0 B/op	       0 allocs/op
BenchmarkWordStructRandom     	       5	 214403633 ns/op	       0 B/op	       0 allocs/op
BenchmarkWordStructRandom     	       5	 214765796 ns/op	       0 B/op	       0 allocs/op
BenchmarkWordStructRandom     	       5	 214246018 ns/op	       0 B/op	       0 allocs/op
BenchmarkWordStructRandom     	       5	 215990274 ns/op	       0 B/op	       0 allocs/op
BenchmarkStruct75Sequential   	      69	  16736356 ns/op	       0 B/op	       0 allocs/op
BenchmarkStruct75Sequential   	      69	  16655312 ns/op	       0 B/op	       0 allocs/op
BenchmarkStruct75Sequential   	      69	  16557392 ns/op	       0 B/op	       0 allocs/op
BenchmarkStruct75Sequential   	      69	  16629100 ns/op	       0 B/op	       0 allocs/op
BenchmarkStruct75Sequential   	      69	  16587764 ns/op	       0 B/op	       0 allocs/op
BenchmarkStruct75Random       	       5	 248547925 ns/op	       0 B/op	       0 allocs/op
BenchmarkStruct75Random       	       5	 248621949 ns/op	       0 B/op	       0 allocs/op
BenchmarkStruct75Random       	       5	 245981673 ns/op	       0 B/op	       0 allocs/op
BenchmarkStruct75Random       	       5	 245869587 ns/op	       0 B/op	       0 allocs/op
BenchmarkStruct75Random       	       5	 246888494 ns/op	       0 B/op	       0 allocs/op
BenchmarkStruct50Sequential   	      64	  18316490 ns/op	       0 B/op	       0 allocs/op
BenchmarkStruct50Sequential   	      64	  18306790 ns/op	       0 B/op	       0 allocs/op
BenchmarkStruct50Sequential   	      64	  18313628 ns/op	       0 B/op	       0 allocs/op
BenchmarkStruct50Sequential   	      64	  18340017 ns/op	       0 B/op	       0 allocs/op
BenchmarkStruct50Sequential   	      66	  18325070 ns/op	       0 B/op	       0 allocs/op
BenchmarkStruct50Random       	       3	 415501346 ns/op	       0 B/op	       0 allocs/op
BenchmarkStruct50Random       	       3	 417333157 ns/op	       0 B/op	       0 allocs/op
BenchmarkStruct50Random       	       3	 416798971 ns/op	       0 B/op	       0 allocs/op
BenchmarkStruct50Random       	       3	 416187758 ns/op	       0 B/op	       0 allocs/op
BenchmarkStruct50Random       	       3	 416526692 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/pavelzagorodnyuk/experiments/cpu_cache/types	73.519s
```
