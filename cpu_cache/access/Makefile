BENCH := ^(BenchmarkSequentialAccess|BenchmarkRandomAccess)$$

.PHONY: bench

bench:
	go test -run='^$$' -bench='$(BENCH)' -benchmem -count=5 -cpu=1
