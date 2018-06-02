# heapsort

A naive heapsort implementation in golang, to test the performance against the standard `sort` package.

## Benchmark

```
$ go test -bench=. -benchmem
```

Output:

```bash
goos: darwin
goarch: amd64
pkg: github.com/alextanhongpin/go-heapsort
BenchmarkHeapSort-4          	20000000	        98.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkSortInt-4           	10000000	       168 ns/op	      80 B/op	       2 allocs/op
BenchmarkHeapSortItem-4      	20000000	        97.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkSortSliceStable-4   	 3000000	       420 ns/op	     256 B/op	       4 allocs/op
PASS
```