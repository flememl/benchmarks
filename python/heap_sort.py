#!/usr/bin/python
from native_sort import unsorted

def _swap(tab, i, j):
    tab[i], tab[j] = tab[j], tab[i]

def _heapify(tab):
    start = (len(tab) - 2) / 2
    while start >= 0:
        _sift_down(tab, start, len(tab)-1)
        start -= 1

def _sift_down(tab, start, end):
    root = start
    while root * 2 + 1 <= end:
        child = root * 2 + 1
        swap = root
        if tab[swap] < tab[child]:
            swap = child
        if child + 1 <= end and tab[swap] < tab[child+1]:
            swap = child + 1
        if swap != root:
            _swap(tab, root, swap)
            root = swap
        else:
            return

def heap_sort(tab):
    tab2 = tab[:]
    _heapify(tab2)
    end = len(tab2)-1
    while end > 0:
        _swap(tab2, end, 0)
        end -= 1
        _sift_down(tab2, 0, end)
    return tab2

heap_sort(unsorted)
