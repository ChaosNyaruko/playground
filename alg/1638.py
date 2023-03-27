"""
给你两个字符串 s 和 t ，请你找出 s 中的非空子串的数目，这些子串满足替换 一个不同字符 以后，是 t 串的子串。换言之，请你找到 s 和 t 串中 恰好 只有一个字符不同的子字符串对的数目。

比方说， "computer" and "computation" 只有一个字符不同： 'e'/'a' ，所以这一对子字符串会给答案加 1 。

请你返回满足上述条件的不同子字符串对数目。

一个 子字符串 是一个字符串中连续的字符。
"""
class Solution:
    def countSubstrings(self, s: str, t: str) -> int:
        m = len(s)
        n = len(t)
        dpl = [[0 for _ in range(n+1)] for _ in range(m+1)]
        dpr = [[0 for _ in range(n+1)] for _ in range(m+1)]
        for i in range(m):
            for j in range(n):
                if (s[i] == t[j]):
                    dpl[i+1][j+1] = dpl[i][j] + 1
                else:
                    dpl[i+1][j+1] = 0
                
                if (s[m -1 - i] == t[n - 1 - j]):
                    dpr[m - i - 1][n - j - 1] = dpr[m - i][n - j] + 1
                else:
                    dpr[m - i - 1][n - j - 1] = 0
        res = 0
        for i in range(m):
            for j in range(n):
                if s[i] != t[j]:
                    res += (dpl[i][j] +1)* (dpr[i+1][j+1]+1)
        
        return res
