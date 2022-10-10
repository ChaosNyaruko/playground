def generate(n):
    res = []
    def dfs(left, right, cur):
        if left == n and right == n:
            res.append(cur)
            return
        if left > n:
            return

        dfs(left+1, right, cur + '(')
        if left > right:
            dfs(left, right+1, cur+')')
        return

    dfs(1, 0, '(')
    return res



if __name__ == '__main__':
    res = generate(1)
    print(res)

'''
Intuition

To enumerate something, generally we would like to express it as a sum of disjoint subsets that are easier to count.

Consider the closure number of a valid parentheses sequence S: the least index >= 0 so that S[0], S[1], ..., S[2*index+1] is valid. Clearly, every parentheses sequence has a unique closure number. We can try to enumerate them individually.

Algorithm

For each closure number c, we know the starting and ending brackets must be at index 0 and 2*c + 1. Then, the 2*c elements between must be a valid sequence, plus the rest of the elements must be a valid sequence.
'''
