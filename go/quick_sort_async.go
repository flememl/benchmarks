package main

import "math/rand"

func QuickSortAsync(tab []int, c chan []int) {
	if len(tab) <= 1 {
		c <- tab
		return
	}
	m := len(tab) / 2
	p := tab[m]
	tab1, tab2 := make([]int, 0), make([]int, 0)
	c1, c2 := make(chan []int), make(chan []int)
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
	unsorted := rand.Perm(1000)
	c := make(chan []int)
	go QuickSortAsync(unsorted, c)
	<-c
}
