```
$ go test -bench .
goos: linux
goarch: amd64
pkg: benchmark
cpu: Intel(R) Core(TM) i9-7900X CPU @ 3.30GHz
BenchmarkEmpty-20               11422674                92.72 ns/op
Benchmark10kb-20                  381254              2764 ns/op
Benchmark100kb-20                  49155             26213 ns/op
Benchmark1000kb-20                  3561            322170 ns/op
Benchmark10x10kb-20                45339             25988 ns/op
Benchmark100x10kb-20                3918            288278 ns/op
PASS
ok      benchmark       7.609s
```
