""" Extra Questions for Lab 13 """

# Q8
def make_generators_generator(g):
    """Generates all the "sub"-generators of the generator returned by
    the generator function g.

    >>> def ints_to(n):
    ...     for i in range(1, n + 1):
    ...          yield i
    ...
    >>> def ints_to_5():
    ...     for item in ints_to(5):
    ...         yield item
    ...
    >>> for gen in make_generators_generator(ints_to_5):
    ...     print("Next Generator:")
    ...     for item in gen:
    ...         print(item)
    ...
    Next Generator:
    1
    Next Generator:
    1
    2
    Next Generator:
    1
    2
    3
    Next Generator:
    1
    2
    3
    4
    Next Generator:
    1
    2
    3
    4
    5
    """
    "*** YOUR CODE HERE ***"
    def f(elem):
      yield elem
    elem = []
    for sg in g():
      elem.append(sg)
      # print(sg)
      yield from f(elem)

