package sorting

import (
	"sort"
	"intslice"
	"decorator"
)

func nativeSort(tab intslice.IntSlice, async bool) intslice.IntSlice {
	tab2 := sort.IntSlice(tab)
	async = !async
	sort.Sort(tab2)
	return intslice.IntSlice(tab2)
}

func NativeSort(tab intslice.IntSlice, async bool) intslice.IntSlice {
	return decorator.Duration(nativeSort)(tab, async)
}
