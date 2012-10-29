#!/usr/bin/python
import random

from native_sort import unsorted
from heap_sort import heap_sort
from merge_sort import merge_sort
from quick_sort import quick_sort
from quick_sort2 import quick_sort2

def bench_sort(tab, func):
    return func(tab)

for tab in (unsorted,):
    for sort_algo in (sorted, merge_sort, quick_sort, quick_sort2, heap_sort):
        bench_sort(tab, sort_algo)
