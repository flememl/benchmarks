package sorting

import (
	"intslice"
	"decorator"
)

func quickSortSync(tab intslice.IntSlice) intslice.IntSlice {
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
	return append(append(quickSortSync(tab1), p), quickSortSync(tab2)...)
}

func quickSortAsync(tab intslice.IntSlice, c chan intslice.IntSlice) {
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
	go quickSortAsync(tab1, c1)
	go quickSortAsync(tab2, c2)
	tab1, tab2 = <-c1, <-c2
	c <- append(append(tab1, p), tab2...)
}

func quickSort(tab intslice.IntSlice, async bool) intslice.IntSlice {
	if async == true {
		c := make(chan intslice.IntSlice)
		go quickSortAsync(tab, c)
		return <-c
	}
	return quickSortSync(tab)
}

func QuickSort(tab intslice.IntSlice, async bool) intslice.IntSlice {
	return decorator.Duration(quickSort)(tab, async)
}
