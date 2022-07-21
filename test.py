from operator import add
def curry2(f):
    print("passed to curry:", f)
    def g(x):
        print("passed to curry2.g",x)
        def h(y):
            print("passed to curry2.g.h", y)
            return f(x,y)
        return h
    return g
