package main

import (
	"flag"
	"intslice"
	"sorting"
)

func main() {
	async := flag.Bool("async", false, "launch asynchronous sorting algos")
	sync := flag.Bool("sync", false, "launch synchronous sorting algos")
	native := flag.Bool("native", false, "launch the native sort")
	merge := flag.Bool("merge", false, "launch the merge sort")
	quick := flag.Bool("quick", false, "launch the quick sort")
	quick2 := flag.Bool("quick2", false, "launch the quick sort in-place")
	flag.Parse()
	sorts := map[*bool] func (intslice.IntSlice, bool) intslice.IntSlice {
	        native: sorting.NativeSort,
		merge: sorting.MergeSort,
		quick: sorting.QuickSort,
		quick2: sorting.QuickSort2,
	}
	run_all := !(*native || *merge || *quick || *quick2)
	for b, f := range sorts {
		if run_all || *b {
			unsorted := intslice.Perm(1000)
			f(unsorted, *async)
			if (*async == true && *sync == true) {
				f(unsorted, false)
			}
		}
	}
}
