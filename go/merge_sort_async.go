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

func MergeSortAsync(tab intslice.IntSlice, c chan intslice.IntSlice) {
	if l := len(tab); l > 1 {
		m := l / 2
		c1, c2 := make(chan intslice.IntSlice), make(chan intslice.IntSlice)
		go MergeSortAsync(tab[:m], c1)
		go MergeSortAsync(tab[m:], c2)
		c <- mergeTables(<-c1, <-c2)
	}
	c <- tab
}

func main() {
	unsorted, c := intslice.Perm(1000), make(chan intslice.IntSlice)
	go MergeSortAsync(unsorted, c)
	<-c
}
