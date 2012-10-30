package intslice

import "math/rand"

type IntSlice []int

func (t IntSlice) Len() (x int) {
	return len(t)
}

func (t IntSlice) Less(i, j int) bool {
	return t[i] < t[j]
}

func (t *IntSlice) Pop() (x int) {
	x, *t = (*t)[0], (*t)[1:]
	return
}

func (t IntSlice) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func Perm(x int) (s IntSlice) {
	return rand.Perm(x)
}
