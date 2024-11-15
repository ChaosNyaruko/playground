import re
from functools import lru_cache
from itertools import product
from collections import deque

class Solution:
    def findMinStep(self, board: str, hand: str) -> int:
        ans = self.dfs(board, "".join(sorted(hand)))
        return ans if ans <= 5 else -1

    @lru_cache(None)
    def dfs(self, cur_board: str, cur_hand: str):
        if not cur_board:
            return 0

        res = 6
        for i, j in product(range(len(cur_board) + 1), range(len(cur_hand))):
            # 第 1 个剪枝条件: 手中颜色相同的球只需要考虑其中一个即可
            if j > 0 and cur_hand[j] == cur_hand[j - 1]:
                continue

            # 第 2 个剪枝条件: 只在连续相同颜色的球的开头位置插入新球
            if i > 0 and cur_board[i - 1] == cur_hand[j]:
                continue

            # 第 3 个剪枝条件: 只考虑放置新球后有可能得到更优解的位置
            #  - 第 1 种情况 : 当前球颜色与后面的球的颜色相同
            #  - 第 2 种情况 : 当前后颜色相同且与当前颜色不同时候放置球
            choose = False
            if 0 < i < len(cur_board) and cur_board[i - 1] == cur_board[i] and cur_board[i - 1] != cur_hand[j]:
                choose = True
            if i < len(cur_board) and cur_board[i] == cur_hand[j]:
                choose = True
            if choose:
                new_board = self.clean(cur_board[:i] + cur_hand[j] + cur_board[i:])
                new_hand = cur_hand[:j] + cur_hand[j + 1:]
                res = min(res, self.dfs(new_board, new_hand) + 1)
        return res

    @staticmethod
    def clean(s):
        n = 1
        while n:
            s, n = re.subn(r'(.)\1{2,}', '', s)
        return s



class Solution:
    def findMinStep(self, board: str, hand: str) -> int:
        def clean(s):
            # 消除桌面上需要消除的球
            n = 1
            while n:
                s, n = re.subn(r"(.)\1{2,}", "", s)
            return s

        hand = "".join(sorted(hand))

        # 初始化用队列维护的状态队列：其中的三个元素分别为桌面球状态、手中球状态和回合数
        queue = deque([(board, hand, 0)])

        # 初始化用哈希集合维护的已访问过的状态
        visited = {(board, hand)}

        while queue:
            cur_board, cur_hand, step = queue.popleft()
            for i, j in product(range(len(cur_board) + 1), range(len(cur_hand))):
                # 第 1 个剪枝条件: 当前球的颜色和上一个球的颜色相同
                if j > 0 and cur_hand[j] == cur_hand[j - 1]:
                    continue

                # 第 2 个剪枝条件: 只在连续相同颜色的球的开头位置插入新球
                if i > 0 and cur_board[i - 1] == cur_hand[j]:
                    continue

                # 第 3 个剪枝条件: 只在以下两种情况放置新球
                #  - 第 1 种情况 : 当前球颜色与后面的球的颜色相同
                #  - 第 2 种情况 : 当前后颜色相同且与当前颜色不同时候放置球
                choose = False
                if 0 < i < len(cur_board) and cur_board[i - 1] == cur_board[i] and cur_board[i - 1] != cur_hand[j]:
                    choose = True
                if i < len(cur_board) and cur_board[i] == cur_hand[j]:
                    choose = True

                if choose:
                    new_board = clean(cur_board[:i] + cur_hand[j] + cur_board[i:])
                    new_hand = cur_hand[:j] + cur_hand[j + 1:]
                    if not new_board:
                        return step + 1
                    if (new_board, new_hand) not in visited:
                        queue.append((new_board, new_hand, step + 1))
                        visited.add((new_board, new_hand))

        return -1


