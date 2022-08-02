# README

- 执行测试 `go test -bench=. .`


- 测试结果

```
goos: darwin
goarch: amd64
pkg: gopl.io/ex/ex1.3
cpu: Intel(R) Core(TM) i5-4258U CPU @ 2.40GHz
BenchmarkConcat-4         705594              1577 ns/op
BenchmarkJoin-4          3806314               330.3 ns/op
PASS
ok      gopl.io/ex/ex1.3        3.248s
```
