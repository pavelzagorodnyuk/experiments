BENCH := ^(BenchmarkFalseSharing|BenchmarkNormalOperation)$$

.PHONY: bench

bench:
	go test -run='^$$' -bench='$(BENCH)' -benchmem -count=5
