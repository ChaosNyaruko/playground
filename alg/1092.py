class Solution:
  def shortestCommonSupersequence(self, str1: str, str2: str) -> str:
    n = len(str1)
    m = len(str2)
    dp = [[0 for _ in range(m+1) ]for _ in range(n+1)]
    for j in range(m):
      dp[n][j] = m - j
    for i in range(n):
      dp[i][m] = n - i
    dp[n][m] = 0
    for i in range(n-1, -1, -1):
      for j in range(m-1, -1, -1):
        if str1[i] == str2[j]:
          dp[i][j] = dp[i+1][j+1] + 1
        else:
          dp[i][j] = min(dp[i+1][j], dp[i][j+1]) + 1

    res = ""
    t1, t2 = 0, 0
    while t1 < n and t2 < m:
      if str1[t1] == str2[t2]:
        res += str1[t1]
        t1 += 1
        t2 += 1
      elif dp[t1][t2] == dp[t1+1][t2]+1:
        res += str1[t1]
        t1 += 1
      else:
        res += str2[t2]
        t2 += 1
    if t1 < n:
      res += str1[t1:]
    if t2 < m:
      res += str2[t2:]
    return res

if __name__ == '__main__':
  print(Solution().shortestCommonSupersequence("abac", "cab")) # cabac
  print(Solution().shortestCommonSupersequence("aaaaaaaaa", "aaaaaaaa")) # aaaaaaaaa
