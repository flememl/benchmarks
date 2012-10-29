#!/usr/bin/python
import random

unsorted = [x for x in xrange(1000)]
random.shuffle(unsorted)

sorted(unsorted)
