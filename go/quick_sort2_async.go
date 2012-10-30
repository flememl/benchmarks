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

func QuickSort2Async(tab intslice.IntSlice, c chan bool) {
	if len(tab) > 1 {
		p := partition(tab)
		c1, c2 := make(chan bool), make(chan bool)
		go QuickSort2Async(tab[:p], c1)
		go QuickSort2Async(tab[p+1:], c2)
		<-c1
		<-c2
	}
	c <- true
}

func main() {
	unsorted, c := intslice.Perm(1000), make(chan bool)
	go QuickSort2Async(unsorted, c)
	<-c
}
