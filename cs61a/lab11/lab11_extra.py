""" Optional Questions for Lab 11 """

from lab11 import *

# Q5
def hailstone(n):
    """
    >>> for num in hailstone(10):
    ...     print(num)
    ...
    10
    5
    16
    8
    4
    2
    1
    """
    "*** YOUR CODE HERE ***"
    while n != 1:
        yield n
        if n % 2 == 0:
            n = n // 2
        else:
            n = 3 * n + 1

    yield n

# Q6
def repeated(t, k):
    """Return the first value in iterable T that appears K times in a row.

    >>> s = [3, 2, 1, 2, 1, 4, 4, 5, 5, 5]
    >>> repeated(trap(s, 7), 2)
    4
    >>> repeated(trap(s, 10), 3)
    5
    >>> print(repeated([4, None, None, None], 3))
    None
    """
    assert k > 1
    "*** YOUR CODE HERE ***"
    cnt = 1
    t = iter(t)
    prev = next(t)
    for cur in t:
        if cur == prev:
            cnt += 1
        else:
            cnt = 1
        if cnt == k:
            return cur
        prev = cur
    

# Q7
def merge(s0, s1):
    """Yield the elements of strictly increasing iterables s0 and s1, removing
    repeats. Assume that s0 and s1 have no repeats. s0 or s1 may be infinite
    sequences.

    >>> m = merge([0, 2, 4, 6, 8, 10, 12, 14], [0, 3, 6, 9, 12, 15])
    >>> type(m)
    <class 'generator'>
    >>> list(m)
    [0, 2, 3, 4, 6, 8, 9, 10, 12, 14, 15]
    >>> def big(n):
    ...    k = 0
    ...    while True: yield k; k += n
    >>> m = merge(big(2), big(3))
    >>> [next(m) for _ in range(11)]
    [0, 2, 3, 4, 6, 8, 9, 10, 12, 14, 15]
    """
    i0, i1 = iter(s0), iter(s1)
    e0, e1 = next(i0, None), next(i1, None)
    "*** YOUR CODE HERE ***"
    while e0 is not None and  e1 is not None:
        if e0 < e1:
            yield e0
            e0 = next(i0, None)
        elif e0 > e1:
            yield e1
            e1 = next(i1, None)
        else:
            yield e0 # or e1
            e0 = next(i0, None)
            e1 = next(i1, None)

    while e0 is not None:
        yield e0
        e0 = next(i0, None)

    while e1 is not None:
        yield e1
        e1 = next(i1, None)
# Q8
def remainders_generator(m):
    """
    Takes in an integer m, and yields m different remainder groups
    of m.

    >>> remainders_mod_four = remainders_generator(4)
    >>> for rem_group in remainders_mod_four:
    ...     for _ in range(3):
    ...         print(next(rem_group))
    0
    4
    8
    1
    5
    9
    2
    6
    10
    3
    7
    11
    """
    "*** YOUR CODE HERE ***"
    def g(r):
        i = 0
        while True:
            yield i * m + r
            i += 1
    return [g(r) for r in range(m)]

# Q9
def zip_generator(*iterables):
    """
    Takes in any number of iterables and zips them together.
    Returns a generator that outputs a series of lists, each
    containing the nth items of each iterable.
    >>> z = zip_generator([1, 2, 3], [4, 5, 6], [7, 8])
    >>> for i in z:
    ...     print(i)
    ...
    [1, 4, 7]
    [2, 5, 8]
    """
    "*** YOUR CODE HERE ***"
    iters = [iter(i) for i in iterables]
    while True:
        try:
            yield [next(it) for it in iters]
        except StopIteration as e:
            # print("runtime err {}, stop".format(e))
            return
