#!/usr/bin/python
from native_sort import unsorted

def _merge_tables(tab1, tab2):
    tab = []
    while len(tab1) or len(tab2):
        if len(tab1) and len(tab2):
            if tab1[0] <= tab2[0]:
                tab.append(tab1.pop(0))
            else:
                tab.append(tab2.pop(0))
        elif len(tab1):
            tab.append(tab1.pop(0))
        elif len(tab2):
            tab.append(tab2.pop(0))
    return tab

def merge_sort(tab):
    l = len(tab)
    if l <= 1:
        return tab
    m = l / 2
    return _merge_tables(merge_sort(tab[:m]), merge_sort(tab[m:]))

merge_sort(unsorted)
