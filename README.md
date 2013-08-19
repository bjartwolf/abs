abs
===

Absolute value in Go using Hacker's delight tricks
So far not able to speed things up... 

go test -bench .

These are my benchmarks so far
==============================
On my dell XPS 13 with Ubuntu 12.04, kernel 3.8.something I think
```shell
Sane way
500000000	         3.64 ns/op
Fast way I
500000000	         4.40 ns/op
Fast way II
500000000	         4.40 ns/op
Fast way III
500000000	         5.45 ns/op
Fast way IV
500000000	         4.77 ns/op
```

On my Surface Pro with Windows 8 
```shell
Sane way
2000000000	         1.08 ns/op
Fast way I
2000000000	         1.26 ns/op
Fast way II
2000000000	         1.22 ns/op
Fast way III
2000000000	         1.67 ns/op
Fast way IV
2000000000	         1.24 ns/op
```
