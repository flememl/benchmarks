package main

import (
	"sort"
	"math/rand"
)

func main() {
	var unsorted sort.IntSlice
	unsorted = rand.Perm(1000)
	sort.Sort(unsorted)
}
