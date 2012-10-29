#!/usr/bin/python
from native_sort import unsorted

def _swap(tab, i1, i2):
    tab[i1], tab[i2] = tab[i2], tab[i1]

def _partition(tab, i1, i2, i):
    p = tab[i]
    _swap(tab, i, i2)
    s = i1
    for k, t in enumerate(tab[:-1]):
        if t < p:
            _swap(tab, k, s)
            s += 1
    _swap(tab, s, i2)
    return s

def quick_sort2(tab):
    l = len(tab)
    if l <= 1:
        return tab
    s = _partition(tab, 0, l-1, l / 2)
    return quick_sort2(tab[:s]) + [tab[s]] + quick_sort2(tab[s+1:])

quick_sort2(unsorted)
