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

func QuickSort2Async(tab []int, c chan bool) {
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
	unsorted, c := rand.Perm(1000), make(chan bool)
	go QuickSort2Async(unsorted, c)
	<-c
}
