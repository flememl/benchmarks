package decorator

import (
	"fmt"
	"intslice"
	"time"
	"runtime"
)

type M func (intslice.IntSlice, bool) intslice.IntSlice

func Duration(f M) M {
	return func (tab intslice.IntSlice, async bool) intslice.IntSlice {
		pc, _, _, _ := runtime.Caller(1)
		me := runtime.FuncForPC(pc)
		t_start := time.Now()
		result := f(tab, async)
		fmt.Printf("%-15s: %s\n", me.Name(), time.Since(t_start).String())
		return result
	}
}
