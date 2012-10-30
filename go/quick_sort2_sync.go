package main

import "intslice"

func partition(tab intslice.IntSlice) (i int) {
	l := len(tab)
	mid, last := l/2, l-1
	tab.Swap(mid, last)
	for i, j := 0, 0; j < last; j++ {
		if tab[j] <= tab[last] {
			tab.Swap(i, j)
			i++
		}
	}
	tab.Swap(i, last)
	return
}

func QuickSort2Sync(tab intslice.IntSlice) {
	if len(tab) > 1 {
		p := partition(tab)
		QuickSort2Sync(tab[:p])
		QuickSort2Sync(tab[p+1:])
	}
}

func main() {
	unsorted := intslice.Perm(1000)
	QuickSort2Sync(unsorted)
}
