package sorting

import (
	"intslice"
	"decorator"
)

func mergeTables(tab1, tab2 intslice.IntSlice) (tab intslice.IntSlice) {
	for len(tab1) > 0 || len(tab2) > 0 {
		if len(tab1) > 0 && len(tab2) > 0 {
			if tab1[0] <= tab2[0] {
				tab = append(tab, tab1.Pop())
			} else {
				tab = append(tab, tab2.Pop())
			}
		} else if len(tab1) > 0 {
			tab = append(tab, tab1.Pop())
		} else if len(tab2) > 0 {
			tab = append(tab, tab2.Pop())
		}
	}
	return
}

func mergeSortSync(tab intslice.IntSlice) intslice.IntSlice {
	if l := len(tab); l > 1 {
		m := l / 2
		return mergeTables(mergeSortSync(tab[:m]), mergeSortSync(tab[m:]))
	}
	return tab
}

func mergeSortAsync(tab intslice.IntSlice, c chan intslice.IntSlice) {
	if l := len(tab); l > 1 {
		m := l / 2
		c1, c2 := make(chan intslice.IntSlice), make(chan intslice.IntSlice)
		go mergeSortAsync(tab[:m], c1)
		go mergeSortAsync(tab[m:], c2)
		c <- mergeTables(<-c1, <-c2)
	}
	c <- tab
}

func mergeSort(tab intslice.IntSlice, async bool) intslice.IntSlice {
	if async == true {
		c := make(chan intslice.IntSlice)
		go mergeSortAsync(tab, c)
		return <-c
	}
	return mergeSortSync(tab)
}

func MergeSort(tab intslice.IntSlice, async bool) intslice.IntSlice {
	return decorator.Duration(mergeSort)(tab, async)
}
