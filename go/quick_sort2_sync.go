package main

import "math/rand"

func swap(tab []int, i, j int) {
	tab[i], tab[j] = tab[j], tab[i]
}

func partition(tab []int) (i int) {
	l := len(tab)
	mid, last := l/2, l-1
	swap(tab, mid, last)
	for i, j := 0, 0; j < last; j++ {
		if tab[j] <= tab[last] {
			swap(tab, i, j)
			i++
		}
	}
	swap(tab, i, last)
	return
}

func QuickSort2Sync(tab []int) {
	if len(tab) > 1 {
		p := partition(tab)
		QuickSort2Sync(tab[:p])
		QuickSort2Sync(tab[p+1:])
	}
}

func main() {
	unsorted := rand.Perm(1000)
	QuickSort2Sync(unsorted)
}
