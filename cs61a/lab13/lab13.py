""" Lab 13: Final Review """

# Q3
def permutations(lst):
    """Generates all permutations of sequence LST. Each permutation is a
    list of the elements in LST in a different order.

    The order of the permutations does not matter.

    >>> sorted(permutations([1, 2, 3]))
    [[1, 2, 3], [1, 3, 2], [2, 1, 3], [2, 3, 1], [3, 1, 2], [3, 2, 1]]
    >>> type(permutations([1, 2, 3]))
    <class 'generator'>
    >>> sorted(permutations((10, 20, 30)))
    [[10, 20, 30], [10, 30, 20], [20, 10, 30], [20, 30, 10], [30, 10, 20], [30, 20, 10]]
    >>> sorted(permutations("ab"))
    [['a', 'b'], ['b', 'a']]
    """
    if not lst:
        yield []
        return
    "*** YOUR CODE HERE ***"
    if len(lst) == 1:
        yield lst[:]
        return
    v = lst[0]
    for rest in permutations(lst[1:]):
        # print("rest:", rest)
        for i in range(len(rest) + 1):
            yield list(rest[:i]) + [v] + list(rest[i:])
