package main

import "math/rand"

func pop(t *[]int) (x int) {
	x, *t = (*t)[0], (*t)[1:]
	return
}

func mergeTables(tab1, tab2 []int) (tab []int) {
	for len(tab1) > 0 || len(tab2) > 0 {
		if len(tab1) > 0 && len(tab2) > 0 {
			if tab1[0] <= tab2[0] {
				tab = append(tab, pop(&tab1))
			} else {
				tab = append(tab, pop(&tab2))
			}
		} else if len(tab1) > 0 {
			tab = append(tab, pop(&tab1))
		} else if len(tab2) > 0 {
			tab = append(tab, pop(&tab2))
		}
	}
	return
}

func MergeSortAsync(tab []int, c chan []int) {
	if l := len(tab); l > 1 {
		m := l / 2
		c1, c2 := make(chan []int), make(chan []int)
		go MergeSortAsync(tab[:m], c1)
		go MergeSortAsync(tab[m:], c2)
		c <- mergeTables(<-c1, <-c2)
	}
	c <- tab
}

func main() {
	unsorted, c := rand.Perm(1000), make(chan []int)
	go MergeSortAsync(unsorted, c)
	<-c
}
