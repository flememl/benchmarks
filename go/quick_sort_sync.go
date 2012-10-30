package main

import "intslice"

func QuickSortSync(tab intslice.IntSlice) intslice.IntSlice {
	if len(tab) <= 1 {
		return tab
	}
	m := len(tab) / 2
	p := (tab)[m]
	tab1, tab2 := make(intslice.IntSlice, 0), make(intslice.IntSlice, 0)
	for i, t := range tab {
		if i == m {
			continue
		}
		if t <= p {
			tab1 = append(tab1, t)
		} else {
			tab2 = append(tab2, t)
		}
	}
	return append(append(QuickSortSync(tab1), p), QuickSortSync(tab2)...)
}

func main() {
	unsorted := intslice.Perm(1000)
	QuickSortSync(unsorted)
}
