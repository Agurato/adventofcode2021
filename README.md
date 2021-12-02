# Advent of Code 2021

## Benchmarks

These are the benchmarks I ran with my inputs

### Day 1

```
goos: windows
goarch: amd64
pkg: vmonot.dev/aoc2021
cpu: Intel(R) Core(TM) i5-8400 CPU @ 2.80GHz
BenchmarkDay1Part1-6   	  973014	      1124 ns/op	       0 B/op	       0 allocs/op
BenchmarkDay1Part2-6   	  202838	      5482 ns/op	   16384 B/op	       1 allocs/op
PASS
ok  	vmonot.dev/aoc2021	2.507s
```

### Day 2

```
goos: windows
goarch: amd64
pkg: vmonot.dev/aoc2021
cpu: Intel(R) Core(TM) i5-8400 CPU @ 2.80GHz
BenchmarkDay2Part1-6   	   15399	     76214 ns/op	   32002 B/op	    1000 allocs/op
BenchmarkDay2Part2-6   	   15667	     76881 ns/op	   32002 B/op	    1000 allocs/op
PASS
ok  	vmonot.dev/aoc2021	4.170s
```