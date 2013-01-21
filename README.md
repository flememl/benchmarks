benchmarks
==========

python/go benchmarks

usage
-----

`./bench_sorts [opts]`

common options:
* no parameter: run all the sorting algorithms
* `--native`: run the default sort
* `--merge`: run mergesort
* `--quick`: run quicksort
* `--quick2`: run quicksort with in-place replacement

python:
* `--heap`: run heapsort

go:
* `--async`: run the algorithms asynchronously
* `--sync`: run the algorithms synchronously (default behaviour)

python
------

profiling: run `python -m cProfile xxx.py` (more info can be found on http://docs.python.org/2/library/profile.html)

go
--

Do not forget to set $GOPATH and to install the packages :
`go install intslice`
`go install sorting`

profiling: not tested yet (more info can be found on http://blog.golang.org/2011/06/profiling-go-programs.html)

you can still run `go build xxx.go` then `time ./xxx`
