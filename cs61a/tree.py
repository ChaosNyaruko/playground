class Tree:
    def __init__(self, label, branches = []):
        self.label = label
        for branch in branches:
            assert isinstance(branch, Tree)
        self.branches = list(branches)

    def __repr__(self):
        if self.branches:
            branch_str = ', ' + repr(self.branches)
        else:
            branch_str = ''

        return 'Tree({0}{1})'.format(self.label, branch_str)

    def __str__(self):
        return '\n'.join(self.indented())

    def indented(self, k = 0):
        indented = []
        for b in self.branches:
            for line in b.indented(k+1):
                indented.append('  ' + line)
        return [str(self.label)] + indented

    def is_leaf(self):
        return not self.branches

def memo(f):
    cache = {}
    def memoized(n):
        if n not in cache:
            cache[n] = f(n)
        return cache[n]
    return memoized

def fib_tree_nomemo(n):
    """A Fibonacci tree. Using `python3 -m doctest test.py` to test it

    >>> print(fib_tree(4))
    3
      1
        0
        1
      2
        1
        1
          0
          1
    """
    if n == 0 or n == 1:
        return Tree(n)
    else:
        left = fib_tree_nomemo(n - 2)
        right = fib_tree_nomemo(n-1)
        fib_n = left.label + right.label
        return Tree(fib_n, [left, right])

@memo
def fib_tree(n):
    """A Fibonacci tree.

    >>> print(fib_tree(4))
    3
      1
        0
        1
      2
        1
        1
         0
         1
    """
    if n == 0 or n == 1:
        return Tree(n)
    else:
        left = fib_tree(n - 2)
        right = fib_tree(n-1)
        fib_n = left.label + right.label
        return Tree(fib_n, [left, right])

def prune_repeats(t, seen):
    t.branches = [b for b in t.branches if b not in seen]
    seen.append(t)
    for b in t.branches:
        prune_repeats(b, seen)

