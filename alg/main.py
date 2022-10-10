from collections import deque
from typing import List

import random

class Trie:
    def __init__(self):
        self.children = [None] * 26
        self.isWord = True
    def insert(self, w):
        p = self
        for c in w:
            i = ord(c) - ord('a')
            if not p.children[i]:
                p.children[i] = Trie()
            p = p.children[i]
        p.isWord = True

    def dfs(self, w, start) -> bool:
        if start == len(w):
            return True
        p = self
        for i in range(start, len(w)):
            c = ord(w[i]) - ord('a')
            if not p.children[c]:
                return False
            p = p.children[c]
            if p.isWord and self.dfs(w, i+1):
                return True
        return False


class Solution:
    def findAllConcatenatedWordsInADict(self, words: List[str]) -> List[str]:
        ans = []
        words.sort(key=len)
        root = Trie()
        for word in words:
            if root.dfs(word, 0):
                ans.append(word)
            else:
                root.insert(word)
        return ans

    def subsetsWithDup(self, nums: List[int]) -> List[List[int]]:
        res = []
        nums.sort()
        def dfs(start, path):
            res.append(path.copy())

            for i in range(start, len(nums)):
                if i > start and nums[i] == nums[i - 1]:
                    continue
                path.append(nums[i])
                dfs(i + 1, path)
                path.pop()

            return
        path = []
        dfs(0, path)
        return res

    def minOperations(self, target: List[int], arr: List[int]) -> int:
        pos = dict()
        for i, t in enumerate(target):
            pos[t] = i

        d = []
        for v in arr:
            if v in pos:
                idx = pos[v]

                # LIS
                if not d or idx > d[-1]:
                    d.append(idx)
                else:
                    l, r = 0, len(d) - 1
                    # find the first element that >= idx, and replace it with idx
                    while l < r:
                        m = l + (r - l) // 2
                        if d[m] < idx:
                            l = m + 1
                        else:
                            r = m
                    d[l] = idx

        return len(target) - len(d)


    def fourSum(self, nums: List[int], target: int) -> List[List[int]]:
        if len(nums) < 4:
            return []

        res = []
        nums.sort()
        n = len(nums)
        for i in range(0, n - 3):
            if i > 0 and nums[i] == nums[i - 1]: continue
            min1 = nums[i] + nums[i + 1] + nums[i + 2] + nums[i + 3]
            if min1 > target:
                break
            max1 = nums[i] + nums[n - 1] + nums[n - 2] + nums[n - 3]
            if max1 < target:
                continue
            for j in range(i + 1, n - 2):
                if j > i + 1 and nums[j] == nums[j - 1]: continue
                min2 = nums[i] + nums[j] + nums[j + 1] + nums[j + 2]
                if min2 > target:
                    break
                max2 = nums[i] + nums[j] + nums[n - 1] + nums[n - 2]
                if max2 < target:
                    continue

                l, r = j + 1, n - 1
                t = target - nums[i] - nums[j]
                while l < r:
                    cur = nums[l] + nums[r]
                    if cur < t:
                        l += 1
                    elif cur == t:
                        res.append([nums[i], nums[j], nums[l], nums[r]])
                        l += 1
                        while l < r and nums[l] == nums[l - 1]:
                            l += 1
                        r -= 1
                        while l < r and nums[r] == nums[r + 1]:
                            r -= 1
                    else:
                        r -= 1

        return res

    def outerTrees(self, trees: List[List[int]]) -> List[List[int]]:
        points = sorted(trees)

        def cross(o, a, b):
            x = [a[0] - o[0], a[1] - o[1]]
            y = [b[0] - o[0], b[1] - o[1]]

            return x[0] * y[1] - x[1] * y[0]

        lower = []
        for p in points:
            while len(lower) >= 2 and cross(lower[-2], lower[-1], p) < 0:
                lower.pop()
            lower.append(p)

        upper = []
        for p in reversed(points):
            while len(upper) >= 2 and cross(upper[-2], upper[-1], p) < 0:
                upper.pop()
            upper.append(p)



        lower = [(p[0], p[1]) for p in lower]
        upper = [(p[0], p[1]) for p in upper]
        # return lower + upper
        return list(set(lower + upper))
    def isCounsins(self, root, x, y):
        if not root:
            return False

        q = deque([root])
        xExist, yExist = False, False
        while q:
            for _ in range(len(q), 0, -1):
                cur = q.popleft()
                if cur.val == x:
                    xExist = True
                if cur.val == y:
                    yExist = True
                l, r = cur.left, cur.right
                if l and r:
                    if l.val == x and r.val == y or r.val == x and l.val == y:
                        return False
                if l:
                    q.append(l)
                if r:
                    q.append(r)

            if xExist != yExist:
                return False
            if xExist and yExist:
                return True

        return False

    def nextGreaterElement(self, nums1: List[int], nums2: List[int]) -> List[int]:
        stk = []
        m = dict()
        for num in nums2:
            while stk and stk[-1] < num:
                m[stk.pop()] = num
            stk.append(num)

        for i in range(len(nums1)):
            if nums1[i] in m:
                nums1[i] = m[nums1[i]]
            else:
                nums1[i] = -1
        return nums1

    def reverseWords(self, s: str) -> str:
        def reverse(s, b, e):
            i, j = b, e
            while i < j:
                s[i], s[j] = s[j], s[i]
                i += 1
                j -= 1

        def reverseWord(s):
            i, j = 0, 0
            n = len(s)
            while i < n:
                while i < j or i < n and s[i] == ' ': i += 1
                while j < i or j < n and s[j] != ' ': j += 1
                reverse(s, i, j - 1)

        def cleanSpaces(s):
            # remove leading and trailing spaces
            # only one space between two words
            i, j = 0, 0
            n = len(s)
            while j < n:
                while j < n and s[j] == ' ': j += 1
                while j < n and s[j] != ' ':
                    s[i] = s[j]
                    i += 1
                    j += 1
                while j < n and s[j] == ' ': j += 1
                if j < n:
                    s[i] = ' '
                    i += 1
            return s[:i]



        s = list(s)
        reverse(s, 0, len(s) - 1)
        reverseWord(s)
        return ''.join(cleanSpaces(s))

    def sortColors(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        ones = -1
        twos = len(nums)
        i =0
        # 1 1 2 0 0 1
        while i < twos:
            print(i, ones, twos)
            if nums[i] == 0:
                ones += 1
                nums[i], nums[ones] = nums[ones], nums[i]
                i += 1
            elif nums[i] == 2:
                twos -= 1
                nums[i], nums[twos] = nums[twos], nums[i]
            else: # nums[i] == 1
                i += 1
            print(nums)

    def orangesRotting(self, grid: List[List[int]]) -> int:
        freshCount = 0
        q = deque()
        m, n = len(grid), len(grid[0])
        for i in range(m):
            for j in range(n):
                if grid[i][j] == 1:
                    freshCount += 1
                elif grid[i][j] == 2:
                    q.append((i, j))

        lvl = 0
        while q:
            for _ in range(len(q), 0, -1):
                cur = q.popleft()
                for d in [(-1, 0), (1,0), (0, 1), (0, -1)]:
                    nx, ny = cur[0] + d[0], cur[1] + d[1]
                    if nx >= m or nx < 0 or ny >= n or ny < 0:
                        continue
                    if grid[nx][ny] == 1:
                        grid[nx][ny] = 2
                        q.append((nx, ny))
                        freshCount -= 1
            lvl += 1

        return lvl - 1 if freshCount == 0 else -1

        def solve(self, board: List[List[str]]) -> None:
            """
                Do not return anything, modify board in-place instead.
            """
            m, n = len(board), len(board[0])
            def dfs(x, y):
                board[nx][ny] = '-'
                for dx, dy in [(-1, 0), (1, 0), (0, -1), (0, 1)]:
                    nx, dy = x + dx, y+dy
                    if nx < 0 or nx >= m or nx <0 or nx >= n:
                        continue
                    if board[nx][ny] == 'O':
                        dfs(nx, ny)

            for i in range(m):
                if board[i][0] == 'O': dfs(i, 0)
                if board[i][n - 1] == 'O': dfs(i, n - 1)

            for j in range(n):
                if board[0][j] == 'O': dfs(0, j)
                if board[m - 1][j] == 'O': dfs(m-1, j)

            for i in range(m):
                for j in range(n):
                    if board[i][j] == 'O':
                        board[i][j] = 'X'
                    elif board[i][j] == '-':
                        board[i][j] = 'O'

            return
    def kInversePairs(self, n: int, k: int) -> int:
        #f[i][j] numbers arrays  of j InversePairs for 1..i
        #f[i][j] = f[i-1][j]+f[i-1][j-1]+....+f[i-1][j-(i-1)] 1..i-1 at most add i-1 more inversepairs
        #f[i][j-1] = f[i-1][j-1] +f[i-1][j-2]+..+f[i-1][j-1-(i-1)]
        #f[i][j] - f[i][j-1] = f[i-1][j] - f[i-1][j-i]
        mod  = 10**9+7
        dp = [[0] * (k+1) for _ in range(n+1)]
        for i in range(n+1):
            dp[i][0] = 1
        for j in range(1, k+1):
            dp[0][j] = 0
        for i in range(1, n+1):
            for j in range(1, k+1):
                dp[i][j] = (dp[i-1][j] - (dp[i-1][j-i] if i <= j else 0) +dp[i][j-1]) % mod

        # print(dp)
        return dp[n][k]

    def getMoneyAmount(self, n: int) -> int:
        # f[i][j] money amount when the answer is between [i, j]
        # f[1][n] is what we want
        # f[i][j] = min(x + max(f[i, x - 1], f[x+1, j]))
        # corner case:
        #   i == j: 0
        #   i > j: 0
        f = [[float('inf')] * (n+1) for _ in range(n+1)]
        for i in range(n+1):
            for j in range(i, -1, -1):
                f[i][j] = 0

        for l in range(2, n+1):
            for a in range(1, n+2 - l):
                b = a + l - 1 # n+1-l-1+l = n
                ab = float('inf')
                for x in range(a, b+1):
                    if x+1 > n: break
                    ab = min(ab, x + max(f[a][x-1], f[x+1][b]))
                f[a][b] = ab
                print(a, b, ab, l, b - a+1)


        return f[1][n]
    def canPartitionKSubsets(self, nums: List[int], k: int) -> bool:
        s = sum(nums)
        if s % k != 0:
            return False
        l = s // k
        v = [0] * len(nums)
        self.flag = False
        def dfs(now, cnt):
            if self.flag:
                return True
            if cnt == k:
                self.flag = True
            if now == l:
                dfs(0, cnt+1)
            for i in range(len(nums)):
                if v[i] == 0:
                    v[i] = 1
                    dfs(now+nums[i], cnt)
                    v[i] = 0
            return
        dfs(0, 0)
        return self.flag


class Cards:

    def __init__(self, nums: List[int]):
        self.ori = nums.copy()
        self.nums = nums


    def reset(self) -> List[int]:
        self.nums = self.ori.copy()


    def shuffle(self) -> List[int]:
        for i in range(len(self.nums)-1, -1, -1):
            idx = random.randint(i, len(self.nums) - 1)
            self.nums[idx], self.nums[i] = self.nums[i], self.nums[idx]
    def __str__(self):
        return str(self.nums)



    # Your Solution object will be instantiated and called as such:
    # obj = Solution(nums)
    # param_1 = obj.reset()
    # param_2 = obj.shuffle()
if __name__ == '__main__':
    c = Cards([1,2,4,5,6])
    c.shuffle()
    print(c)
    c.reset()
    print(c)

    sl = Solution()
    # nums = [2,2,2,2,2]
    # print(sl.fourSum(nums, 8))

    # 示例 1：

    # 输入：target = [5,1,3], arr = [9,4,2,3,4]
    # 输出：2
    # 解释：你可以添加 5 和 1 ，使得 arr 变为 [5,9,4,1,2,3,4] ，target 为 arr 的子序列。
    # 示例 2：

    # 输入：target = [6,4,8,1,3,2], arr = [4,7,6,2,3,8,6,1]
    # 输出：3
    # print(sl.minOperations([5,1,3], [9,4,2,3,4]))
    # print(sl.minOperations([6,4,8,1,3,2], [4,7,6,2,3,8,6,1]))
    # print(sl.subsetsWithDup([1,2,2]))
    # print(sl.outerTrees([[1,2],[2,2],[4,2]]))
    # print(sl.outerTrees([[1,1],[2,2],[2,0],[2,4],[3,3],[4,2]]))
    # print(sl.reverseWords(" hello world     "))
    # x = [2,0,2,1,1,0]
    # print(x)
    # sl.sortColors(x)
    # print(x)

    # g = [[2,1,1],[1,1,0],[0,1,1]]
    # print(sl.orangesRotting(g))
    # print(sl.kInversePairs(3, 0))
    # print(sl.getMoneyAmount(100))
    # ans = sl.canPartitionKSubsets([3522,181,521,515,304,123,2512,312,922,407,146,1932,4037,2646,3871,269], 5)

    # ans = sl.findAllConcatenatedWordsInADict(["cat","cats","catsdogcats","dog","dogcatsdog","hippopotamuses","rat","ratcatdogcat"])
    ans = sl.findAllConcatenatedWordsInADict([""])
    print(ans)




