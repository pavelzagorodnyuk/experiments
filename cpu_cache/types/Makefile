BENCH := ^(BenchmarkWordStruct|BenchmarkStruct75|BenchmarkStruct50)(Sequential|Random)$$

.PHONY: bench

bench:
	go test -run='^$$' -bench='$(BENCH)' -benchmem -count=5 -cpu=1
