package main

import "intslice"

func QuickSortAsync(tab intslice.IntSlice, c chan intslice.IntSlice) {
	if len(tab) <= 1 {
		c <- tab
		return
	}
	m := len(tab) / 2
	p := tab[m]
	tab1, tab2 := make(intslice.IntSlice, 0), make(intslice.IntSlice, 0)
	c1, c2 := make(chan intslice.IntSlice), make(chan intslice.IntSlice)
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
	go QuickSortAsync(tab1, c1)
	go QuickSortAsync(tab2, c2)
	tab1, tab2 = <-c1, <-c2
	c <- append(append(tab1, p), tab2...)
}

func main() {
	unsorted, c := intslice.Perm(1000), make(chan intslice.IntSlice)
	go QuickSortAsync(unsorted, c)
	<-c
}
