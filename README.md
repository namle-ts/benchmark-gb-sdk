# benchmark-gb-sdk

## Usage

```bash
 go test -bench=. -count=5 -benchmem
```
Output
```
goos: darwin
goarch: arm64
pkg: github.com/namle-ts/benchmark-gb-sdk
BenchmarkEvaluateFeature-8       1862744               681.0 ns/op          1064 B/op         17 allocs/op
BenchmarkEvaluateFeature-8       1656321               637.0 ns/op          1064 B/op         17 allocs/op
BenchmarkEvaluateFeature-8       1853337               630.3 ns/op          1064 B/op         17 allocs/op
BenchmarkEvaluateFeature-8       1943397               608.3 ns/op          1064 B/op         17 allocs/op
BenchmarkEvaluateFeature-8       1958786               624.1 ns/op          1064 B/op         17 allocs/op
PASS
ok      github.com/namle-ts/benchmark-gb-sdk    9.365s
```
---
```bash
go test -bench=. -count=5 -benchmem -benchtime=5s
```
Output
```
goos: darwin
goarch: arm64
pkg: github.com/namle-ts/benchmark-gb-sdk
BenchmarkEvaluateFeature-8       9529080               639.7 ns/op          1064 B/op         17 allocs/op
BenchmarkEvaluateFeature-8       9664617               624.6 ns/op          1064 B/op         17 allocs/op
BenchmarkEvaluateFeature-8       9686954               612.4 ns/op          1064 B/op         17 allocs/op
BenchmarkEvaluateFeature-8       9758883               636.9 ns/op          1064 B/op         17 allocs/op
BenchmarkEvaluateFeature-8       9817269               609.8 ns/op          1064 B/op         17 allocs/op
PASS
ok      github.com/namle-ts/benchmark-gb-sdk    33.729s
```