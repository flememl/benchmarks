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

func MergeSortSync(tab []int) []int {
	if l := len(tab); l > 1 {
		m := l / 2
		return mergeTables(MergeSortSync(tab[:m]), MergeSortSync(tab[m:]))
	}
	return tab
}

func main() {
	unsorted := rand.Perm(1000)
	MergeSortSync(unsorted)
}
