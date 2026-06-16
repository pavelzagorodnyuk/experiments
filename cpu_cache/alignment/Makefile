BENCH := ^(BenchmarkAlignedStruct|BenchmarkPadded25Struct|BenchmarkPadded50Struct)(Sequential|Random)$$

.PHONY: bench

bench:
	go test -run='^$$' -bench='$(BENCH)' -benchmem -count=5 -cpu=1
