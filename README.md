benchmarks
==========

python/go benchmarks

python
------

sorting algorithms: heap sort, merge sort, quick sort & quick sort with in-place replacement

profiling: run `python -m cProfile xxx.py` (more info can be found on http://docs.python.org/2/library/profile.html)

go
--

sorting algorithms: merge sort, quick sort & quick sort with in-place replacement, both synchronous and asynchronous (mostly to play with goroutines and channels)

Do not forget to set $GOPATH and to install the intslice package with `go install intslice`.

profiling: not tested yet (more info can be found on http://blog.golang.org/2011/06/profiling-go-programs.html)

you can still run `go build xxx.go` then `time ./xxx`
