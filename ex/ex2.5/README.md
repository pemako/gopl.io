# README

```
go test -bench=.
goos: darwin
goarch: amd64
pkg: gopl.io/ex/ex2.5
cpu: Intel(R) Core(TM) i5-4258U CPU @ 2.40GHz
BenchmarkTable-4                1000000000               0.4841 ns/op
BenchmarkTableLoop-4            323081199                4.134 ns/op
BenchmarkShiftValue-4           30484189                33.53 ns/op
BenchmarkClearRightmost-4       100000000               12.74 ns/op
PASS
ok      gopl.io/ex/ex2.5        5.583s
```
