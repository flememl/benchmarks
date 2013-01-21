def quick_sort(tab):
    if len(tab) <= 1:
        return tab
    m = len(tab) / 2
    p = tab[m]
    tab1, tab2 = [], []
    for t in tab[:m] + tab[m+1:]:
        if t <= p:
            tab1.append(t)
        else:
            tab2.append(t)
    return quick_sort(tab1) + [p] + quick_sort(tab2)
