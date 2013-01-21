package sorting

import (
	"intslice"
	"decorator"
)

func partition(tab intslice.IntSlice) (i int) {
	l := len(tab)
	mid, last := l/2, l-1
	tab.Swap(mid, last)
	for j := 0; j < last; j++ {
		if tab[j] <= tab[last] {
			tab.Swap(i, j)
			i++
		}
	}
	tab.Swap(i, last)
	return
}

func quickSort2Sync(tab intslice.IntSlice) {
	if len(tab) > 1 {
		p := partition(tab)
		quickSort2Sync(tab[:p])
		quickSort2Sync(tab[p+1:])
	}
}

func quickSort2Async(tab intslice.IntSlice, c chan bool) {
	if len(tab) > 1 {
		p := partition(tab)
		c1, c2 := make(chan bool), make(chan bool)
		go quickSort2Async(tab[:p], c1)
		go quickSort2Async(tab[p+1:], c2)
		<-c1
		<-c2
	}
	c <- true
}

func quickSort2(tab intslice.IntSlice, async bool) intslice.IntSlice {
	if async == true {
		c := make(chan bool)
		go quickSort2Async(tab, c)
		<-c
	} else {
		quickSort2Sync(tab)
	}
	return tab
}

func QuickSort2(tab intslice.IntSlice, async bool) intslice.IntSlice {
	return decorator.Duration(quickSort2)(tab, async)
}
