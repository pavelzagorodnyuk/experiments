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
pkg: github.com/pavelzagorodnyuk/experiments/cpu_cache/alignment
cpu: AMD Ryzen 7 9700X 8-Core Processor             
BenchmarkAlignedStructSequential  	      50	  23738826 ns/op	       0 B/op	       0 allocs/op
BenchmarkAlignedStructSequential  	      50	  23641156 ns/op	       0 B/op	       0 allocs/op
BenchmarkAlignedStructSequential  	      50	  23609523 ns/op	       0 B/op	       0 allocs/op
BenchmarkAlignedStructSequential  	      50	  23692278 ns/op	       0 B/op	       0 allocs/op
BenchmarkAlignedStructSequential  	      50	  23642963 ns/op	       0 B/op	       0 allocs/op
BenchmarkAlignedStructRandom      	       5	 236881933 ns/op	       0 B/op	       0 allocs/op
BenchmarkAlignedStructRandom      	       5	 236989697 ns/op	       0 B/op	       0 allocs/op
BenchmarkAlignedStructRandom      	       5	 235983336 ns/op	       0 B/op	       0 allocs/op
BenchmarkAlignedStructRandom      	       5	 236357560 ns/op	       0 B/op	       0 allocs/op
BenchmarkAlignedStructRandom      	       5	 236555911 ns/op	       0 B/op	       0 allocs/op
BenchmarkPadded25StructSequential 	      51	  23288567 ns/op	       0 B/op	       0 allocs/op
BenchmarkPadded25StructSequential 	      51	  23324160 ns/op	       0 B/op	       0 allocs/op
BenchmarkPadded25StructSequential 	      51	  23329524 ns/op	       0 B/op	       0 allocs/op
BenchmarkPadded25StructSequential 	      51	  23315578 ns/op	       0 B/op	       0 allocs/op
BenchmarkPadded25StructSequential 	      51	  23365896 ns/op	       0 B/op	       0 allocs/op
BenchmarkPadded25StructRandom     	       6	 182758351 ns/op	       0 B/op	       0 allocs/op
BenchmarkPadded25StructRandom     	       6	 182838655 ns/op	       0 B/op	       0 allocs/op
BenchmarkPadded25StructRandom     	       6	 183462851 ns/op	       0 B/op	       0 allocs/op
BenchmarkPadded25StructRandom     	       6	 182847904 ns/op	       0 B/op	       0 allocs/op
BenchmarkPadded25StructRandom     	       6	 182788035 ns/op	       0 B/op	       0 allocs/op
BenchmarkPadded50StructSequential 	      49	  24504011 ns/op	       0 B/op	       0 allocs/op
BenchmarkPadded50StructSequential 	      49	  24459180 ns/op	       0 B/op	       0 allocs/op
BenchmarkPadded50StructSequential 	      49	  24509210 ns/op	       0 B/op	       0 allocs/op
BenchmarkPadded50StructSequential 	      49	  24460995 ns/op	       0 B/op	       0 allocs/op
BenchmarkPadded50StructSequential 	      49	  24501170 ns/op	       0 B/op	       0 allocs/op
BenchmarkPadded50StructRandom     	       9	 120554806 ns/op	       0 B/op	       0 allocs/op
BenchmarkPadded50StructRandom     	       9	 120330747 ns/op	       0 B/op	       0 allocs/op
BenchmarkPadded50StructRandom     	       9	 120503123 ns/op	       0 B/op	       0 allocs/op
BenchmarkPadded50StructRandom     	       9	 120602222 ns/op	       0 B/op	       0 allocs/op
BenchmarkPadded50StructRandom     	       9	 120254113 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/pavelzagorodnyuk/experiments/cpu_cache/alignment	54.196s
```
