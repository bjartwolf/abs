abs
===

Absolute value in Go using Hacker's delight tricks
So far not able to speed things up... 

go test -bench .

These are my benchmarks so far
==============================
On my dell XPS 13 with Ubuntu 12.04, kernel 3.8.something I think
...Outdated

On my Surface Pro with Windows 8 
It seems 8 bit numbers are faster with method 4 (involving multiplication) but the fastest in general in method 1.

```shell
PASS
BenchmarkSaneAbs8	50000000	        28.2 ns/op
BenchmarkInSaneAbs8	100000000	        21.2 ns/op
BenchmarkInSaneAbs82	100000000	        21.2 ns/op
BenchmarkInSaneAbs83	100000000	        22.9 ns/op
BenchmarkInSaneAbs84	100000000	        18.1 ns/op
BenchmarkSaneAbs	100000000	        26.7 ns/op
BenchmarkInSaneAbs	100000000	        21.2 ns/op
BenchmarkInSaneAbs2	100000000	        21.7 ns/op
BenchmarkInSaneAbs3	100000000	        25.1 ns/op
BenchmarkInSaneAbs4	100000000	        24.2 ns/op
ok  	github.com/bjartwolf/abs	22.028s
```
