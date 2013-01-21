package main

import "intslice"

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

func MergeSortSync(tab intslice.IntSlice) intslice.IntSlice {
	if l := len(tab); l > 1 {
		m := l / 2
		return mergeTables(MergeSortSync(tab[:m]), MergeSortSync(tab[m:]))
	}
	return tab
}

func main() {
	unsorted := intslice.Perm(1000)
	MergeSortSync(unsorted)
}
