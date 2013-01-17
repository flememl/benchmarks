import time

def timing(func):
    def wrapper(tab, sort):
        t_start = time.time()
        func(tab, sort)
        print "%-15s: %f" % (sort.__name__, time.time() - t_start)
    return wrapper
