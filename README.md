# slice-sum

This project demonstrates some hardware and software peculiarities
arising during simple task of calculating sum of slice elements.

Let's take `int64` slice, fill it with random values and find sum of all elements.
But instead of just iterate over slice and adding total (`Classic`), we try something that called
[vectorization](https://en.wikipedia.org/wiki/Single_instruction,_multiple_data)

Run on Apple M2 Pro, 32GB

```
go test -bench=. -test.benchtime=10s -test.cpu=1                             
goos: darwin
goarch: arm64
pkg: slice-sum
BenchmarkClassicSum                          411          29113861 ns/op
BenchmarkVectorized1Sum                      414          28852641 ns/op
BenchmarkVectorized2Sum                      830          14469991 ns/op
BenchmarkVectorized4Sum                      763          15824555 ns/op
BenchmarkVectorized4Sum2Acc                 1648           7269878 ns/op
BenchmarkVectorized8Sum                      794          15055297 ns/op
BenchmarkVectorized8Sum2Acc                  804          14848083 ns/op
BenchmarkVectorized16Sum                     554          21862634 ns/op
BenchmarkVectorized16Sum2Acc                 804          14910646 ns/op
PASS
ok      slice-sum       138.917s
```

`BenchmarkVectorized2Sum` is "faster" then `BenchmarkClassicSum` by factor of 2.
It's not surprising if we consider slice element size (64 bit)
and [cache line size](https://en.wikipedia.org/wiki/CPU_cache#Cache_entries):

```
sysctl -a hw machdep.cpu | grep cachelinesize
hw.cachelinesize: 128
```

More surprising for me was that `BenchmarkVectorized4Sum2Acc` is faster then `BenchmarkClassicSum` by factor of 4,
and also faster by `BenchmarkVectorized4Sum` by factor of 2!

So I dag it deep and with help of profiler found out that there is no magic there.
It's just [inlining](https://en.wikipedia.org/wiki/Inline_expansion).

If we put `go:noinline` pragma before `BenchmarkVectorized4Sum2Acc` it's astonishing speed disappears:

```
BenchmarkClassicSum                          411          29113861 ns/op
...
BenchmarkVectorized4Sum2Acc                 1648           7269878 ns/op
BenchmarkVectorized4Sum2AccNoInline          676          17780127 ns/op
...
```

With this simple example we can see that our simple daily tasks
still require knowledge of hardware and software features to be solved efficiently.
